package main

import (
	// "AppFactory/internal/pkg/rsa/security"
	"AppFactory/pkg/rsa/security"
	"fmt"
	"io/ioutil"
)

func main() {

	//testEncDecPem()
	//testEncDecPfx()
	//testSignPem()
	//testSignPfx()
	aes_key := "ed4f8731b6ae7a19"
	encry_msg := "k4pBEkQrvic457iQ/xaL5n3dm/vH+atfYRfSCxHsymg="

	data := AesDecry(aes_key, encry_msg)
	fmt.Printf("data[%s]\n", data)

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
	signbuf := "accessType=0&bizType=000201&currencyCode=156"
	signature, err := SignPem([]byte(signbuf), "./etc/local_pri_sys.pem")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("signature[%s]\n", signature)

	err = VertifyPem(signbuf, signature, "./etc/local_pub_sys.pem")
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
	s, err := rsaKey.Sign(security.SHA256WithRSA, signBuf)
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

	return rsaKey.Verify(security.SHA256WithRSA, []byte(signBuf), bs)
}
