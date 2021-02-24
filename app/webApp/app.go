package webapp

import (
	"AppFactory/internal/service"
	"AppFactory/pkg/config"
	"AppFactory/pkg/log"
	"net/http"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// MiddlewareAuth 认证
func MiddlewareAuth(r *ghttp.Request) {
	token := r.Get("token")
	if token == "123456" {
		r.Middleware.Next()
	} else {
		r.Response.WriteStatus(http.StatusForbidden)
	}
}

// MiddlewareCORS 跨域
func MiddlewareCORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

// MiddlewareLog 全局中间件打印日志
func MiddlewareLog(r *ghttp.Request) {

	logger := log.GetLogInstance()
	r.Middleware.Next()
	errStr := ""
	if err := r.GetError(); err != nil {
		errStr = err.Error()
	}
	logger.Info(r.Response.Status, r.URL.Path, errStr)
	if r.Response.Status >= http.StatusInternalServerError {
		r.Response.ClearBuffer()
		r.Response.Writeln("哎哟我去，服务器居然开小差了，请稍后再试吧！")
	}
}

// WebRouterGroup 分组路由
func WebRouterGroup(group *ghttp.RouterGroup) {
	group.Middleware(
		service.Middleware.Ctx,
		service.Middleware.CORS,
	)

	group.Group("/api.v2", func(group *ghttp.RouterGroup) {

		group.Middleware(MiddlewareAuth, MiddlewareCORS)

		group.GET("/test", func(r *ghttp.Request) {
			r.Response.Write("receive it, ok")
			// panic("service err ,hahahah")
		})

	})

}

// RunExec 程序运行主逻辑
func RunExec() {
	s := g.Server()
	// s.Domain("127.0.0.1").BindHandler("/", hello1)
	// s.Domain("localhost").BindHandler("/{class}-{course}/:name/*act", hello2)

	s.Use(MiddlewareLog)

	s.Group("/", WebRouterGroup)
	s.SetPort(g.Cfg().GetInt("application.port"))
	s.Run()
}

// Setup 配置文件加载和组件预先初始化
func Setup(confName string) {
	cfg := config.InitConfig(confName)
	// 日志初始化
	log.InitLogger(cfg)

}
