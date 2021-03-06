package main

import (
	// "AppFactory/internal/pkg/rsa/security"
	tools "AppFactory/pkg"
	"AppFactory/pkg/rsa/security"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	// msg := "admin123"
	// encStr := RsaEncry(msg, `D:\codeRepo\Golang\AppFactory\pkg\rsa\etc\rsa_public_key_copy.pem`)
	// fmt.Printf("加密后：[%s]\n", encStr)
	// text := RsaDecry(encStr, `D:\codeRepo\Golang\AppFactory\pkg\rsa\etc\rsa_private_key_copy.pem`)
	// fmt.Printf("解密后：[%s]\n", text)

	// testSignPem()

	fmt.Println("**************")
	AesCipherStr := `6F34D9A8BF9B30E70B7381B68304752F8BA5BF02BDFA64814E5F65AACBBC003FD482A772F3471E5AAEA32BF8B777719753542DDBA135284B95D24C83B3B6DD3BD9C408E3B896269C287F767757A1C9A3A5AE0AFCE0EE11AB4DC960FB496DD45CAC3A31E710C803BDF910686E234FCF86129BB22348944CC62FA05A047D327527115A9452EC63104670833478A82C43AE8AD71D18F0AD31067EFB26BA418130C315871FE52329EAC998BB9C753174CFBD11855A4800D0D1DB0F9D480672D9E36BECCC3DF50207F1B3BC68AA67B5901B154960AEC9F7D155AAC043FC41E2E6B7F37847831FED4263B59E5DFE917C7FF865FFA5484242A8E0240249CACA9E53E12B`
	sysAesKey := `54NrlJk13wU6pn==`
	byteCipher, _ := hex.DecodeString(AesCipherStr)
	data, err := security.AESDecrypt(byteCipher, []byte(sysAesKey), security.AES_ECB_PKCS5PADDING)
	if err != nil {
		fmt.Printf("用户信息解密失败[%s]", err)
	}
	fmt.Printf("AES解密结果[%s]\n", data)
	fmt.Println("**************")
	// 模拟解析响应报文信息
	formStr := `sign=GUB%2BJQt%2B3A9lU57C6wv%2FUGiW1Yrtdnsw%2FXR7NP5bIPoPG3%2B1obWBY121dJ6N%2BHryCVUDM%2FE3HmIq%2FD42wyZGfkm4QFpEB7ESSq9JQJ9OGWQQFqz72DN%2BTh8O6nWibaIddVZXonI%2FhsdYvyBOqEPjI2IOQ5zN5LGhvdVdMNdFGKE%3D&version=1.0&request_id=2021012711272051719770&content=EA108F24A7ABCA35EC3D0E31F3139B5BF9229E4A8047AEA2E750486C7D7E695077651F8014305BC92980A71C127993476DA6E93E0F07487E471A4CEFD44A7281FB31546A09CB9FDCDC634100589653C6376CDA1E29107E89F3F27765A31EF18420C04534890273ABD9B7859350AF333DC12C298C8A6CB265D0E6C9538901BD2CD64DA6EE66795F242F7A12722A125D3B82D552683A7BD43596FA0D9B0145D0F80AA82A25A515A171EB7BDA2CA8BDE29654995424469320DAE48579890557146F3C779332365535453DCDC4D9F031BA5D09939BF90539BE6CBF25E767D88367941D34A87C8C49D007FFBB2B842D5EA91A7058427CFD6E0431041407F4510EB39056044B5A2CBCBEAB9F97732DBFA488C9022891FF10FCE943189472C32EA17B093EF6EEB532169254E095A70AC8725224FFA5484242A8E0240249CACA9E53E12B`
	item := tools.ParseUrl(formStr)

	// // 输出全部内容
	// for key, value := range item {
	// 	if key == "content" {
	// 		byteCipher, _ := hex.DecodeString(value[0])
	// 		data, _ := security.AESDecrypt([]byte(byteCipher), []byte(sysAesKey), security.AES_ECB_PKCS5PADDING)
	// 		fmt.Println(key, "-->", string(data))
	// 	} else {
	// 		fmt.Println(key, "-->", value[0])
	// 	}

	// }

	type SGNotifyReq struct {
		Version   string `json:"version,omitempty" post:"version,omitempty"`       //版本号
		Sign      string `json:"sign,omitempty" post:"sign,omitempty"`             //签名
		Content   string `json:"content,omitempty" post:"content,omitempty"`       // 业务请求参数集合
		RequestID string `json:"request_id,omitempty" post:"request_id,omitempty"` // 请求流水号
	}
	type SGNotifyContent struct {
		OrderID      string `json:"order_id,omitempty" post:"order_id,omitempty"`
		TxnStatus    string `json:"txn_status,omitempty" post:"txn_status,omitempty"`
		CreateTime   string `json:"create_time,omitempty" post:"create_time,omitempty"`
		ChnlRespCode string `json:"chnl_resp_code,omitempty" post:"chnl_resp_code,omitempty"`
		ChnlRespMsg  string `json:"chnl_resp_msg,omitempty" post:"chnl_resp_msg,omitempty"`
		MerID        string `json:"mer_id,omitempty"  post:"mer_id,omitempty"`            //商户号
		MerOrderID   string `json:"mer_order_id,omitempty" post:"mer_order_id,omitempty"` // 商户订单号
		TxnAmount    string `json:"txn_amount,omitempty" post:"txn_amount,omitempty"`     // 交易金额(分)
		TxnFee       string `json:"txn_fee,omitempty" post:"txn_fee,omitempty"`           // 手续费(分)
		AcctNo       string `json:"acct_no,omitempty" post:"acct_no,omitempty"`           // 账号
		AcctName     string `json:"acct_name,omitempty" post:"acct_name,omitempty"`       // 账号名称 对公账户必填
		GpayMerID    string `json:"gpay_mer_id,omitempty" post:"gpay_mer_id,omitempty"`   // 代付商户号
	}
	// 按需获取
	notifyReq := new(SGNotifyReq)
	notifyContent := new(SGNotifyContent)

	notifyReq.Version = item.Get("version")
	notifyReq.RequestID = item.Get("request_id")
	notifyReq.Sign = item.Get("sign")
	notifyReq.Content = item.Get("content")

	CipherContent, _ := hex.DecodeString(notifyReq.Content)
	ContentJSONStr, _ := security.AESDecrypt([]byte(CipherContent), []byte(sysAesKey), security.AES_ECB_PKCS5PADDING)

	err = VertifyPem(tools.Signbuf2(notifyReq), item.Get("sign"), `D:\codeRepo\Golang\AppFactory\pkg\rsa\etc\rsa_public_key_copy.pem`)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("验签成功")
	}

	json.Unmarshal([]byte(ContentJSONStr), notifyContent)
	fmt.Printf("notifyReq:[%+v]\ncontent:[%+v]\n", notifyReq, notifyContent)

}

func testEncDecPem() {
	msg := "hellp"
	encryMsg := RsaEncry(msg, "./etc/local_pub_sys.pem")
	fmt.Printf("encryMsg[%s]\n", encryMsg)
	decryMsg := RsaDecry(encryMsg, "./etc/local_pri_sys.pem")
	fmt.Printf("decryMsg[%s]\n", decryMsg)
}

func testEncDecPfx() {
	msg2 := "你好"
	encMsg2 := RsaEncryCer(msg2, "./etc/local_pub_sys.cer")
	fmt.Printf("encMsg2[%s]\n", encMsg2)
	decryMsg2 := RsaDecryPfx(encMsg2, "./etc/local_pri_sys.pfx", "cfca1234")
	fmt.Printf("decryMsg2[%s]\n", decryMsg2)
}

func testSignPem() {
	// signature, err := SignPem([]byte(signbuf), `D:\codeRepo\Golang\AppFactory\pkg\rsa\etc\rsa_private_key_copy.pem`)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Printf("signature[%s]\n", signature)
	signbuf := "brh_id=1099991049&content=8BD0EDDAF1ED371DA4D454811723B814A4EF73760321EBE5189652B1409C0C2ADAB1F3B788CC614DF4D6C794EDA95AD17B82CACAF3C9313ADF72585284FBC4E5A389344AFE73FE75840B05F7B71A1233ED54F908C7C696ECB254A68B76874DC76FC61EC56D54C6BA7BD5A4C498F33803E80F3740E8E2BB4CC046CC1A6A99692F7D04F4EA5C76ED50535C3B9899247A303BA0388D2ADB71DC35BE9E2E8805ED7B08F9002D1569575F3C2B28E068A9A25F3CFE720997B158F9F1CC6C20C1838D909B6B7F5435BE8BF3A3B8075F80FCC2DD&request_id=1611207121175&version=1.0"
	Jsignature := "OlYQCTHB5eKYkInk8YZvRxNMAuVM5x3T65A0HVwuL5vOW5JsbaBw0hiTkKujU7foJT1IuIEBDAxNYSzjHuL0Y9+E91N7Kwj7k/el/14Wp+WSUw/ziXidvrVjUP/WtDE3kzvKJM/2ylYcWN8DxzO9Quo9/DPPxcMPedwv+hgSFys="
	// SignedStr := `MAdIYkHPA0cHkIjVS6sbrLEHpf3KXjbv71vqhG3vhRYodhyc1oo1H936MM3gHZOjI0ccusa6WC87GOWYjA2ycQLxCaFlyQtAeMObYnvNEKqLvfW0slbDXzE7163mSH4Nj3ov3bb27SEbB3O85QjDjlNdf2XBGusV8IHF8nJEdx4=`
	signature, err := SignPem([]byte(signbuf), `D:\codeRepo\Golang\AppFactory\pkg\rsa\etc\Jrsa_private_key_copy.pem`)
	fmt.Printf("Jsignature[%s]\n", Jsignature)
	fmt.Printf("Gsignature[%s]\n", signature)
	err = VertifyPem(signbuf, signature, `D:\codeRepo\Golang\AppFactory\pkg\rsa\etc\rsa_public_key_copy.pem`)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("验签成功")
	}
}

func testSignPfx() {
	signbuf := "accessType=0&bizType=000201&currencyCode=156"
	signature, err := SignPfx([]byte(signbuf), "./etc/local_pri_sys.pfx", "cfca1234")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("signature[%s]\n", signature)

	err = VertifyCer(signbuf, signature, "./etc/local_pub_sys.cer")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("验签成功")
	}
}

func RsaDecry(msg, priKey string) string {
	decode64 := security.DecodeBase64([]byte(msg))
	pk, err := ioutil.ReadFile(priKey)
	if err != nil {
		fmt.Printf("读取文件失败[%s]\n", err)
	}
	priK, err := security.LoadRSAPrivatePEM(string(pk))
	if err != nil {
		fmt.Printf("加载私钥失败\n")
	}

	text, err := priK.DecryptPKCS1v15(decode64)
	if err != nil {
		fmt.Println("解密失败[%s]", err)
	}
	return string(text)
}

func RsaEncry(msg, pubKey string) string {
	pk, err := ioutil.ReadFile(pubKey)
	if err != nil {
		fmt.Printf("读取文件失败[%s]\n", err)
	}
	pubK, err := security.LoadRSAPubicPEM(string(pk))
	if err != nil {
		fmt.Printf("加载公钥失败\n")
	}
	cipher, err := pubK.EncryptPKCS1v15([]byte(msg))
	if err != nil {
		fmt.Println("加密失败[%s]", err)
	}
	encode64 := security.EncodeBase64([]byte(cipher))
	return string(encode64)
}

func RsaEncryCer(msg, pubKey string) string {

	cerData, err := ioutil.ReadFile(pubKey)
	if err != nil {
		fmt.Printf("读取文件失败[%s]\n", err)
	}
	pubcert, err := security.NewCertInfoCer(string(cerData))
	if err != nil {
		fmt.Printf("加载公钥证书失败[%s]", err)
	}

	text, err := pubcert.EncryptPKCS1v15([]byte(msg))
	if err != nil {
		fmt.Println("加密失败[%s]", err)
	}
	return string(security.EncodeBase64(text))
}

func RsaDecryPfx(msg, priKey, passwd string) string {
	decode64 := security.DecodeBase64([]byte(msg))
	pfxData, err := ioutil.ReadFile(priKey)
	if err != nil {
		fmt.Printf("读取文件失败[%s]\n", err)
	}
	cert, err := security.NewCertInfoPfx(pfxData, passwd)
	if err != nil {
		fmt.Printf("加载私钥证书失败[%s]", err)
	}
	cipher, err := cert.DecryptPKCS1v15(decode64)
	if err != nil {
		fmt.Println("解密失败[%s]", err)
	}
	return string(cipher)
}

func AesDecry(sysAesKey, encrypted string) string {
	plain := security.DecodeBase64([]byte(encrypted))
	if len(plain) == 0 {
		fmt.Printf("用户信息解密结果为空\n")
	}
	data, err := security.AESDecrypt(plain, []byte(sysAesKey), security.AES_ECB_PKCS5PADDING)
	if err != nil {
		fmt.Printf("用户信息解密失败[%s]", err)
	}
	return string(data)
}

func AesEncry(key, origData string) string {
	cipher, err := security.AESEncrypt([]byte(origData), []byte(key), security.AES_ECB_PKCS5PADDING)

	if err != nil {
		fmt.Println("解密失败[%s]", err)
	}
	return string(security.EncodeBase64(cipher))
}

func SignPfx(signBuf []byte, priKey, passwd string) (string, error) {

	pfxData, err := ioutil.ReadFile(priKey)
	if err != nil {
		fmt.Printf("读取文件失败[%s]\n", err)
	}
	cert, err := security.NewCertInfoPfx(pfxData, passwd)
	if err != nil {
		fmt.Printf("加载私钥证书失败[%s]", err)
	}

	h, _ := security.Hash(security.SHA256WithRSA, signBuf)
	hexstr := security.EncodeHex(h)
	s, err := cert.Sign(security.SHA256WithRSA, hexstr)
	if err != nil {
		return "", err
	}
	return string(security.EncodeBase64(s)), nil
}

func VertifyCer(signBuf, signature, pubKey string) error {
	cerData, err := ioutil.ReadFile(pubKey)
	if err != nil {
		fmt.Printf("读取文件失败[%s]\n", err)
	}
	cert, err := security.NewCertInfoCer(string(cerData))
	if err != nil {
		fmt.Printf("加载公钥证书失败[%s]", err)
	}

	if signature == "" {
		return nil
	}
	bs := security.DecodeBase64([]byte(signature))
	h, err := security.Hash(security.SHA256WithRSA, []byte(signBuf))
	if err != nil {
		return err
	}
	hexstr := security.EncodeHex(h)
	return cert.Verify(security.SHA256WithRSA, hexstr, bs)
}

func SignPem(signBuf []byte, priKey string) (string, error) {
	pk, err := ioutil.ReadFile(priKey)
	if err != nil {
		fmt.Printf("读取文件失败[%s]\n", err)
	}
	rsaKey, err := security.LoadRSAPrivatePEM(string(pk))
	if err != nil {
		fmt.Printf("加载私钥失败\n")
	}
	// 这里修改了sha256WithRsa 为 MD5WithRSA方便测试
	s, err := rsaKey.Sign(security.MD5WithRSA, signBuf)
	if err != nil {
		fmt.Printf("签名失败[%s]", err)
		return "", err
	}
	signature := security.EncodeBase64(s)
	return string(signature), nil
}

func VertifyPem(signBuf, signature, pubKey string) error {
	cerData, err := ioutil.ReadFile(pubKey)
	if err != nil {
		fmt.Printf("读取文件失败[%s]\n", err)
	}
	rsaKey, err := security.LoadRSAPubicPEM(string(cerData))
	if err != nil {
		fmt.Printf("加载公钥证书失败[%s]", err)
	}

	if signature == "" {
		return nil
	}
	bs := security.DecodeBase64([]byte(signature))

	// return rsaKey.Verify(security.SHA256WithRSA, []byte(signBuf), bs)
	return rsaKey.Verify(security.MD5WithRSA, []byte(signBuf), bs)
}
