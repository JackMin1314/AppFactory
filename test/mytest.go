package main

import (
	tools "AppFactory/pkg"
	"fmt"
	"strings"

	// "fmt"

	"encoding/xml"

	"net/http"

	"github.com/axgle/mahonia"
)

// Request
type TLTHktRequest struct {
	XMLName  xml.Name        `xml:"AIPG"`
	Info     *HktPayInfo     `xml:"INFO"`
	Transfer *HktPayTransfer `xml:"TRANSFER"`
}

type HktPayInfo struct {
	XMLName   xml.Name `xml:"INFO"`
	TrxCode   string   `xml:"TRX_CODE"`
	Version   string   `xml:"VERSION"`
	DataType  string   `xml:"DATA_TYPE"`
	Level     string   `xml:"LEVEL"`
	UserName  string   `xml:"USERNAME"`
	UserPass  string   `xml:"USERPASS"`
	ReqSN     string   `xml:"REQ_SN"`
	SignedMsg string   `xml:"SIGNED_MSG"`
	// add for response
	RetCode string `xml:"RET_CODE"`
	ErrMsg  string `xml:"ERR_MSG"`
}

type HktPayTransfer struct {
	XMLName       xml.Name `xml:"TRANSFER"`
	MerchantID    string   `xml:"MERCHANT_ID"`
	SubmitTime    string   `xml:"SUBMIT_TIME"`
	AccountName   string   `xml:"ACCOUNT_NAME"`
	Amount        string   `xml:"AMOUNT"`
	Currency      string   `xml:"CURRENCY"`
	BusinessCode  string   `xml:"BUSINESS_CODE"`
	FromAccountNo string   `xml:"FROM_ACCOUNT_NO"`
	ToAccountNo   string   `xml:"TO_ACCOUNT_NO"`
	Remark        string   `xml:"REMAKR"`
}

// Response
type TLTHktPayResponse struct {
	XMLName  xml.Name           `post:"AIPG" xml:"AIPG"`
	Info     *HktPayRspInfo     `post:"INFO" xml:"INFO"`
	Transfer *HktPayRspTransret `post:"TRANSRET" xml:"TRANSRET,omitempty"`
	//MerID string	`post:"-"`
}

type HktPayRspInfo struct {
	XMLName xml.Name `xml:"INFO"`
	// 必填
	TrxCode   string `xml:"TRX_CODE"`
	Version   string `xml:"VERSION"`
	DataType  string `xml:"DATA_TYPE"`
	ReqSN     string `xml:"REQ_SN"`
	RetCode   string `xml:"RET_CODE"` // 返回错误代码
	ErrMsg    string `xml:"ERR_MSG"`  // 错误信息
	SignedMsg string `xml:"SIGNED_MSG"`
}

type HktPayRspTransret struct {
	XMLName xml.Name `xml:"TRANSRET"`
	// 必填
	RetCode   string `xml:"RET_CODE"`             // 返回码
	ErrMsg    string `xml:"ERR_MSG"`              // 错误文本
	SettleDay string `xml:"SETTLE_DAY,omitempty"` // 清算日期
	// 选填
	VoucherNo string `xml:"VOUCHERNO,omitempty"` // 银行流水
}

func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}
func main() {

	// dataString := `419d9c5f8e95a374ec3d14d65fa0199b384c0731b3c0dbe43927ef7f05132676e701396ea251978a93b8636f0857e95e5d1a62ded9a6831d01249bbd9ecd997da5461a707ea49de54c60512343c09303af932207dd880a362d751795d0b088df81002991119f64cf8ccbfa8a9fefd70858d74f436232c7cf7eac904c3f10b119e46aa4c8604e01eb95775e45900c319ed4e8304ffeef5e20ed1ada3c7d8b54c6786cd5ab9231731d8d01e100cde78b87645926e450e5c70e5bacce8f5ed56e3c0857b693cd4abf44ec20acbdece153f7cd33c0593771cd7026298f0ce0e86f20a723c09c7c3ccacdc95c64133d6a269c482cd15f4f5c0ee1bae2e585ac97496f`
	/**
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Post(``,
		"application/xml",
		bytes.NewBuffer([]byte(dataString)))
	if err != nil {
		fmt.Println("get err: ", err)
	}
	defer resp.Body.Close()
	respData, err := ioutil.ReadAll(resp.Body)

	fmt.Println("get resp xml string:\n", string(respData))

	decoder := xml.NewDecoder(bytes.NewReader(respData))
	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		return tools.GBKReader(charset, input)
	}
	HKTReq := new(TLTHktRequest)
	if err := decoder.Decode(HKTReq); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%+s\n", HKTReq)
	}
	**/

	// GBK to UTF-8
	// strtemp := strings.ReplaceAll(dataString, `<?xml version="1.0" encoding="gbk" ?>`, `<?xml version="1.0" encoding="UTF-8" ?>`)
	// strtemp = strings.ReplaceAll(dataString, `<?xml version="1.0" encoding="GBK" ?>`, `<?xml version="1.0" encoding="UTF-8" ?>`)
	// utfData := tools.GBKtUTF8([]byte(strtemp))
	// fmt.Println("gbk to utf-8\n",string(utfData))

	// fmt.Printf("after decode[%s]", HKTReq.Info.ErrMsg)
	// mistr, err := xml.MarshalIndent(HKTReq,"","	")

	// 写一个服务接受请求
	// http.HandleFunc("/gbkxml", XmlObj)
	// http.ListenAndServe("127.0.0.1:9988", nil)

	dataStr := `<?xml version="1.0" encoding="GBK"?>
	<AIPG>
		<INFO>
			<TRX_CODE>100007</TRX_CODE>
			<VERSION>04</VERSION>
			<DATA_TYPE>2</DATA_TYPE>
			<REQ_SN>200604000000445-0001616656945917</REQ_SN>
			<RET_CODE>1001</RET_CODE>
			<ERR_MSG>提交失败：明细商户号与头部商户号不一致</ERR_MSG>
			<SIGNED_MSG>88e4c450e07c9a005316ac2f04a09222347452818b7a6eed8e051d50c2fc941f3dff36d8cdc1094e3489acf564272e3be2e8f79c24d7c0f156db654cb8165d33316ac9ea7a97d8bea2c540010c6860326c5038306d790d66d0198bf73ea92217a929bd44d9d40c365fef428a8f414916c039fd6dbc1779a4f55fde042b5c63277f105fd5c7e10b4b40715cb92188753778cc807062d01b5311120e4b8943d0345d1c669ed579c96c0ce5e0e841f7b4fcc6880e51260082d709489502d1c3ddd0841d866f4f36fcabeaf43b6f920edf5a833d68cee2c69fea7fe31432ff6d2ee815d75c83e08cfa5bf9f308b8c18c108932d0209b97163da4d4a694ec24d21ff4</SIGNED_MSG>
		</INFO>
	</AIPG>`
	preStr := `<SIGNED_MSG>`
	sufStr := `</SIGNED_MSG>`
	lp := strings.Index(dataStr,preStr)
	rp := strings.Index(dataStr,sufStr)
	fmt.Println(dataStr[:lp]+dataStr[rp+13:])

}

func XmlObj(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "hello world")
	ret := &TLTHktPayResponse{}
	infoItem := new(HktPayRspInfo)
	ret.Info = infoItem

	ret.Info.TrxCode = "100007"
	ret.Info.Version = "04"
	ret.Info.DataType = "2"
	ret.Info.ReqSN = "2021032264710678899853953"
	ret.Info.RetCode = "1000"
	// ret.Info.ErrMsg = "商户用户名不能为空"
	ret.Info.ErrMsg = "签名不符"

	// xml.Unmarshal(bufstruct,ret)
	// respData, _ := xml.MarshalIndent(ret, "", "	")

	w.Header().Set("Content-Type", "application/xml;charset=GBK")
	// w.Write([]byte(`<?xml version="1.0" encoding="GBK"?>` + "\n" + bufstruct))
	bt := tools.XMLGBKEncoder(ret)
	// xml.Unmarshal(bt, ret)
	// respData, _ := xml.MarshalIndent(ret, "", "	")
	w.Write([]byte(`<?xml version="1.0" encoding="GBK"?>` + "\n"+ string(bt)))
}
