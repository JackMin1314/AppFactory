package security

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"

	"AppFactory/pkg/rsa/security/pkcs12"

	"github.com/pkg/errors"
)

type CertInfo struct {
	KeyType      string
	KeyFile      string
	KeyPass      string
	cert         *x509.Certificate
	CertString   string
	serialNumber string
	privateKey   *rsa.PrivateKey
	publicKey    *rsa.PublicKey
}

func NewCertInfo(cfg map[string]string) (*CertInfo, error) {
	var err error
	cert := new(CertInfo)
	cert.KeyFile = cfg["KEY_FILE"]
	cert.KeyType = cfg["KEY_TYPE"]
	cert.KeyPass = cfg["KEY_PASS"] //PFX 密码

	keyBuf, err := ioutil.ReadFile(cert.KeyFile)
	if err != nil {
		return nil, errors.WithMessagef(err, "读取证书文件失败[%s]", cert.KeyFile)
	}

	switch cert.KeyType {
	case FILE_CERT_CER:
		pemBlock, _ := pem.Decode(keyBuf)
		if len(pemBlock.Bytes) == 0 {
			return nil, errors.Errorf("PEM Decode failed[%s]", cert.KeyFile)
		}
		cert.cert, err = x509.ParseCertificate(pemBlock.Bytes)
		if err != nil {
			return nil, errors.Errorf("解析PEM证书[%s]失败[%s]", cert.KeyFile, err)
		}
		if pub, ok := cert.cert.PublicKey.(*rsa.PublicKey); ok {
			cert.publicKey = pub
		} else {
			return nil, errors.Errorf("CERT公钥提取失败[%s]失败", cert.KeyFile)
		}
		cert.serialNumber = cert.cert.SerialNumber.String()
	case FILE_CERT_PFX:
		var pri interface{}
		pri, cert.cert, err = pkcs12.Decode(keyBuf, cert.KeyPass)
		if err != nil {
			return nil, errors.Errorf("PFX证书加载失败[%s]失败[%s]", cert.KeyFile, err)
		}
		if p, ok := pri.(*rsa.PrivateKey); ok {
			cert.privateKey = p
		} else {
			return nil, errors.Errorf("FPX私钥提取失败[%s]失败", cert.KeyFile)
		}
		cert.publicKey = cert.privateKey.Public().(*rsa.PublicKey)
		cert.serialNumber = cert.cert.SerialNumber.String()
	default:
		return nil, errors.Errorf("非法类型[%s]失败", cert.KeyType)
	}
	return cert, nil
}

func NewCertInfoCer(cerkey string) (*CertInfo, error) {
	var err error
	cert := new(CertInfo)

	pemBlock, _ := pem.Decode([]byte(cerkey))
	if len(pemBlock.Bytes) == 0 {
		return nil, errors.Errorf("PEM Decode failed[%s]", cert.KeyFile)
	}
	cert.cert, err = x509.ParseCertificate(pemBlock.Bytes)
	if err != nil {
		return nil, errors.Errorf("解析PEM证书[%s]失败[%s]", cert.KeyFile, err)
	}
	if pub, ok := cert.cert.PublicKey.(*rsa.PublicKey); ok {
		cert.publicKey = pub
	} else {
		return nil, errors.Errorf("CERT公钥提取失败[%s]失败", cert.KeyFile)
	}
	cert.serialNumber = cert.cert.SerialNumber.String()
	cert.CertString = string(cerkey)
	return cert, nil
}

func NewCertInfoCerPString(cerkey string) (*CertInfo, error) {
	var err error
	cert := new(CertInfo)

	pemBlock, _ := pem.Decode([]byte(cerkey))
	if len(pemBlock.Bytes) == 0 {
		return nil, errors.Errorf("PEM Decode failed[%s]", cert.KeyFile)
	}
	cert.cert, err = x509.ParseCertificate(pemBlock.Bytes)
	if err != nil {
		return nil, errors.Errorf("解析PEM证书[%s]失败[%s]", cert.KeyFile, err)
	}
	if pub, ok := cert.cert.PublicKey.(*rsa.PublicKey); ok {
		cert.publicKey = pub
	} else {
		return nil, errors.Errorf("CERT公钥提取失败[%s]失败", cert.KeyFile)
	}
	cert.serialNumber = cert.cert.SerialNumber.String()
	return cert, nil
}

func NewCertInfoPfx(pfxcert []byte, passwd string) (*CertInfo, error) {
	var err error
	cert := new(CertInfo)
	var pri interface{}
	pri, cert.cert, err = pkcs12.Decode(pfxcert, passwd)
	if err != nil {
		return nil, errors.Errorf("PFX证书加载失败[%s]失败[%s]", cert.KeyFile, err)
	}
	if p, ok := pri.(*rsa.PrivateKey); ok {
		cert.privateKey = p
	} else {
		return nil, errors.Errorf("FPX私钥提取失败[%s]失败", cert.KeyFile)
	}
	cert.publicKey = cert.privateKey.Public().(*rsa.PublicKey)
	cert.serialNumber = cert.cert.SerialNumber.String()
	return cert, nil
}

/*algo 算法:
  signbuf 签名串
  signature 签名值
*/
func (cert *CertInfo) Sign(algo int, signbuf []byte) ([]byte, error) {

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
		return nil, x509.InsecureAlgorithmError(algo)
	default:
		return nil, fmt.Errorf("非法的签名算法:[%d]", algo)
	}

	if !hashType.Available() {
		return nil, x509.ErrUnsupportedAlgorithm
	}
	h := hashType.New()
	h.Write(signbuf)
	digest := h.Sum(nil)

	return cert.privateKey.Sign(rand.Reader, digest, hashType)
}

/*algo 算法:
  signbuf 签名串
  signature 签名值
*/
func (cert *CertInfo) Verify(algo int, signbuf []byte, signature []byte) error {
	rsaKey, err := LoadFromRSAPublic(cert.publicKey)
	if err != nil {
		return err
	}
	return rsaKey.Verify(algo, signbuf, signature)
}

/*
	公钥加密:
*/
func (cert *CertInfo) EncryptPKCS1v15(origdata []byte) ([]byte, error) {
	rsaKey, err := LoadFromRSAPublic(cert.publicKey)
	if err != nil {
		return nil, err
	}
	return rsaKey.EncryptPKCS1v15(origdata)
}

/*

	私钥解密
*/
func (cert *CertInfo) DecryptPKCS1v15(cipherdata []byte) ([]byte, error) {
	rsaKey, err := LoadFromRSAPrivate(cert.privateKey)
	if err != nil {
		return nil, err
	}
	return rsaKey.DecryptPKCS1v15(cipherdata)
}

/*获取证书信息*/
func (cert *CertInfo) SerialNumber() string {
	return cert.serialNumber
}
