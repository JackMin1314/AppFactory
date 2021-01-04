package security

import (
	"crypto"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io/ioutil"
)

const (
	//证书RSA
	FILE_CERT_CER = "FILE_CERT_CER"
	FILE_CERT_PFX = "FILE_CERT_PFX"
	//公私钥
	FILE_RSA_PEM_PUB    = "PEM-RSA-PUB"
	FILE_RSA_PEM_PRIV   = "PEM-RSA-PRIV"
	FILE_RSA_PKCS8_PRIV = "FILE_RSA_PKCS8_PRIV"
)

const (
	UnknownSignatureAlgorithm int = iota
	MD2WithRSA
	MD5WithRSA
	SHA1WithRSA
	SHA256WithRSA
	SHA384WithRSA
	SHA512WithRSA
	DSAWithSHA1
	DSAWithSHA256
	ECDSAWithSHA1
	ECDSAWithSHA256
	ECDSAWithSHA384
	ECDSAWithSHA512
)

var algoName = [...]string{
	MD2WithRSA:      "MD2-RSA",
	MD5WithRSA:      "MD5-RSA",
	SHA1WithRSA:     "SHA1-RSA",
	SHA256WithRSA:   "SHA256-RSA",
	SHA384WithRSA:   "SHA384-RSA",
	SHA512WithRSA:   "SHA512-RSA",
	DSAWithSHA1:     "DSA-SHA1",
	DSAWithSHA256:   "DSA-SHA256",
	ECDSAWithSHA1:   "ECDSA-SHA1",
	ECDSAWithSHA256: "ECDSA-SHA256",
	ECDSAWithSHA384: "ECDSA-SHA384",
	ECDSAWithSHA512: "ECDSA-SHA512",
}

type CliCert struct {
	KeyType            string `json:"key_type"`
	KeyFile            string `json:"key_file"`
	SerialNumber       string `json:"serial_number"`
	X509Cert           *x509.Certificate
	PublicKey          interface{}
	PublicKeyAlgorithm x509.PublicKeyAlgorithm
}

func NewCliCert(jsonConfig string) (*CliCert, error) {
	cert := new(CliCert)
	err := json.Unmarshal([]byte(jsonConfig), cert)
	if err != nil {
		return nil, err
	}

	certBuf, err := ioutil.ReadFile(cert.KeyFile)
	if err != nil {
		return nil, fmt.Errorf("密钥文件[%s]读取失败[%s]", cert.KeyFile, err)
	}

	keyBlock, _ := pem.Decode(certBuf)
	if keyBlock == nil {
		return nil, fmt.Errorf("密钥文件[%s]PEM解码失败", cert.KeyFile)
	}

	switch cert.KeyType {
	case FILE_CERT_CER:
		cert.PublicKey, err = x509.ParsePKIXPublicKey(keyBlock.Bytes)
		if err != nil {
			return nil, fmt.Errorf("读取PEM公钥文件失败[%s]", err)
		}
		cert.PublicKeyAlgorithm = x509.RSA
	case FILE_CERT_PFX:
		fallthrough
	default:
		cert.X509Cert, err = x509.ParseCertificate(keyBlock.Bytes)
		if err != nil {
			return nil, fmt.Errorf("解析PEM公钥[%s]失败[%s]", cert.KeyFile, err)
		}
		cert.PublicKey = cert.X509Cert.PublicKey
		cert.PublicKeyAlgorithm = cert.X509Cert.PublicKeyAlgorithm
		cert.SerialNumber = cert.X509Cert.SerialNumber.String()
	}
	return cert, nil
}

/*algo 算法:
  signbuf 签名串
  signature 签名值
*/
func (cert *CliCert) Verfy(algo int, signbuf, signature []byte) error {

	var hashType crypto.Hash

	switch algo {
	case SHA1WithRSA, DSAWithSHA1, ECDSAWithSHA1:
		hashType = crypto.SHA1
	case SHA256WithRSA, DSAWithSHA256, ECDSAWithSHA256:
		hashType = crypto.SHA256
	case SHA384WithRSA, ECDSAWithSHA384:
		hashType = crypto.SHA384
	case SHA512WithRSA, ECDSAWithSHA512:
		hashType = crypto.SHA512
	case MD2WithRSA, MD5WithRSA:
		return x509.InsecureAlgorithmError(algo)
	default:
		return fmt.Errorf("非法的验签算法:[%d]", algo)
	}

	if !hashType.Available() {
		return x509.ErrUnsupportedAlgorithm
	}
	h := hashType.New()
	h.Write(signbuf)
	digest := h.Sum(nil)

	switch pub := cert.PublicKey.(type) {
	case *rsa.PublicKey:
		return rsa.VerifyPKCS1v15(pub, hashType, digest, signature)
	default:
		return x509.ErrUnsupportedAlgorithm
	}
}

/*
	签名公钥加密:
*/
func (cert *CliCert) EncryptPKCS1v15(origdata []byte) ([]byte, error) {
	switch pub := cert.PublicKey.(type) {
	case *rsa.PublicKey:
		rsaKey, err := LoadFromRSAPublic(pub)
		if err != nil {
			return nil, err
		}
		return rsaKey.EncryptPKCS1v15(origdata)
	default:
		return nil, x509.ErrUnsupportedAlgorithm
	}
}

/*获取证书信息*/
func (cert *CliCert) GetSerialNumber() string {
	return cert.SerialNumber
}
