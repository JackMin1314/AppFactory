package api

import (
	"AppFactory/internal/model"
	"AppFactory/internal/service"

	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

// 用户API管理对象
var User = new(userApi)

type userApi struct{}

type JsonResponse struct {
	Code    int         `json:"code"`    // 错误码((0:成功, 1:失败, >1:错误码))
	Message string      `json:"message"` // 提示信息
	Data    interface{} `json:"data"`    // 返回数据(业务接口定义具体数据结构)
}

// 标准返回结果数据结构封装。
func Json(r *ghttp.Request, code int, message string, data ...interface{}) {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}
	r.WriteJson(JsonResponse{
		Code:    code,
		Message: message,
		Data:    responseData,
	})
}

// 返回JSON数据并退出当前HTTP执行函数。
func JsonExit(r *ghttp.Request, err int, msg string, data ...interface{}) {
	Json(r, err, msg, data...)
	r.Exit()
}

// @summary 用户注册接口
// @tags    用户服务
// @produce json
// @param   entity  body model.UserApiSignUpReq true "注册请求"
// @router  /user/signup [POST]
// @success 200 {object} JsonResponse "执行结果"
func (a *userApi) SignUp(r *ghttp.Request) {
	var (
		apiReq     *model.UserApiSignUpReq
		serviceReq *model.UserServiceSignUpReq
	)
	if err := r.ParseForm(&apiReq); err != nil {
		JsonExit(r, 1, err.Error())
	}
	if err := gconv.Struct(apiReq, &serviceReq); err != nil {
		JsonExit(r, 1, err.Error())
	}
	if err := service.User.SignUp(serviceReq); err != nil {
		JsonExit(r, 1, err.Error())
	} else {
		JsonExit(r, 0, "ok")
	}
}

// @summary 用户登录接口
// @tags    用户服务
// @produce json
// @param   passport formData string true "用户账号"
// @param   password formData string true "用户密码"
// @router  /user/signin [POST]
// @success 200 {object} JsonResponse "执行结果"
func (a *userApi) SignIn(r *ghttp.Request) {
	var (
		data *model.UserApiSignInReq
	)
	if err := r.Parse(&data); err != nil {
		JsonExit(r, 1, err.Error())
	}
	if err := service.User.SignIn(r.Context(), data.Passport, data.Password); err != nil {
		JsonExit(r, 1, err.Error())
	} else {
		JsonExit(r, 0, "ok")
	}
}

// @summary 判断用户是否已经登录
// @tags    用户服务
// @produce json
// @router  /user/issignedin [GET]
// @success 200 {object} JsonResponse "执行结果:`true/false`"
func (a *userApi) IsSignedIn(r *ghttp.Request) {
	JsonExit(r, 0, "", service.User.IsSignedIn(r.Context()))
}

// @summary 用户注销/退出接口
// @tags    用户服务
// @produce json
// @router  /user/signout [GET]
// @success 200 {object} JsonResponse "执行结果, 1: 未登录"
func (a *userApi) SignOut(r *ghttp.Request) {
	if err := service.User.SignOut(r.Context()); err != nil {
		JsonExit(r, 1, err.Error())
	}
	JsonExit(r, 0, "ok")
}

// @summary 检测用户账号接口(唯一性校验)
// @tags    用户服务
// @produce json
// @param   passport query string true "用户账号"
// @router  /user/checkpassport [GET]
// @success 200 {object} JsonResponse "执行结果:`true/false`"
func (a *userApi) CheckPassport(r *ghttp.Request) {
	var (
		data *model.UserApiCheckPassportReq
	)
	if err := r.Parse(&data); err != nil {
		JsonExit(r, 1, err.Error())
	}
	if data.Passport != "" && !service.User.CheckPassport(data.Passport) {
		JsonExit(r, 1, "账号已经存在", false)
	}
	JsonExit(r, 0, "", true)
}

// @summary 检测用户昵称接口(唯一性校验)
// @tags    用户服务
// @produce json
// @param   nickname query string true "用户昵称"
// @router  /user/checknickname [GET]
// @success 200 {object} JsonResponse "执行结果"
func (a *userApi) CheckNickName(r *ghttp.Request) {
	var (
		data *model.UserApiCheckNickNameReq
	)
	if err := r.Parse(&data); err != nil {
		JsonExit(r, 1, err.Error())
	}
	if data.Nickname != "" && !service.User.CheckNickName(data.Nickname) {
		JsonExit(r, 1, "昵称已经存在", false)
	}
	JsonExit(r, 0, "ok", true)
}

// @summary 获取用户详情信息
// @tags    用户服务
// @produce json
// @router  /user/profile [GET]
// @success 200 {object} model.User "用户信息"
func (a *userApi) Profile(r *ghttp.Request) {
	JsonExit(r, 0, "", service.User.GetProfile(r.Context()))
}
