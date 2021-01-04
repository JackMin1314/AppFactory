package security

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"math/big"

	"golang.org/x/crypto/pkcs12"
)

type RSAKey struct {
	KeyType    string `json:"key_type"`
	KeyFile    string `json:"key_file"`
	PublicKey  *rsa.PublicKey
	PrivateKey *rsa.PrivateKey
	Modules    int
}

func NewRSA(jsonConfig string) (*RSAKey, error) {
	r := new(RSAKey)
	err := json.Unmarshal([]byte(jsonConfig), r)
	if err != nil {
		return nil, err
	}

	keyBuf, err := ioutil.ReadFile(r.KeyFile)
	if err != nil {
		return nil, fmt.Errorf("读取RSA密钥失败[%s]", r.KeyFile)
	}

	block, _ := pem.Decode(keyBuf)
	if len(block.Bytes) == 0 {
		return nil, fmt.Errorf("读取RSA密钥文件w为空[%s]", r.KeyFile)
	}

	var ok bool
	switch r.KeyType {
	case FILE_RSA_PEM_PUB:
		pub, err := x509.ParsePKIXPublicKey(block.Bytes)
		if err != nil {
			return nil, fmt.Errorf("解析RSA公钥失败[%s]", err)
		}
		r.PublicKey, ok = pub.(*rsa.PublicKey)
		if !ok {
			return nil, fmt.Errorf("Value returned from ParsePKIXPublicKey was not an RSA public key")
		}
		r.Modules = (r.PublicKey.N.BitLen() + 7) / 8
	case FILE_RSA_PEM_PRIV:
		r.PrivateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return nil, fmt.Errorf("ParsePKCS1PrivateKey err:%s", err)
		}
		r.Modules = (r.PrivateKey.N.BitLen() + 7) / 8
	default:
		return nil, fmt.Errorf("非法的RSA公钥类型[%s]", r.KeyType)
	}
	return r, nil
}

func NewRSAInMap(cfg map[string]string) (*RSAKey, error) {
	r := new(RSAKey)

	r.KeyFile = cfg["KEY_FILE"]
	r.KeyType = cfg["KEY_TYPE"]

	keyBuf, err := ioutil.ReadFile(r.KeyFile)
	if err != nil {
		return nil, fmt.Errorf("读取RSA密钥失败[%s]", r.KeyFile)
	}

	block, _ := pem.Decode(keyBuf)
	if len(block.Bytes) == 0 {
		return nil, fmt.Errorf("读取RSA密钥文件w为空[%s]", r.KeyFile)
	}

	var ok bool
	switch r.KeyType {
	case FILE_RSA_PEM_PUB:
		pub, err := x509.ParsePKIXPublicKey(block.Bytes)
		if err != nil {
			return nil, fmt.Errorf("解析RSA公钥失败[%s]", err)
		}
		r.PublicKey, ok = pub.(*rsa.PublicKey)
		if !ok {
			return nil, fmt.Errorf("Value returned from ParsePKIXPublicKey was not an RSA public key")
		}
		r.Modules = (r.PublicKey.N.BitLen() + 7) / 8
	case FILE_RSA_PEM_PRIV:
		r.PrivateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return nil, fmt.Errorf("ParsePKCS1PrivateKey err:%s", err)
		}
		r.Modules = (r.PrivateKey.N.BitLen() + 7) / 8
	case FILE_RSA_PKCS8_PRIV:
		key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			return nil, fmt.Errorf("ParsePKCS1PrivateKey err:%s", err)
		}
		r.PrivateKey = key.(*rsa.PrivateKey)
		r.Modules = (r.PrivateKey.N.BitLen() + 7) / 8
	default:
		return nil, fmt.Errorf("非法的RSA公钥类型[%s]", r.KeyType)
	}
	return r, nil
}

func NewRSAOfCer(cfg map[string]string) (*RSAKey, error) {
	r := new(RSAKey)

	r.KeyFile = cfg["KEY_FILE"]
	r.KeyType = cfg["KEY_TYPE"]

	keyBuf, err := ioutil.ReadFile(r.KeyFile)
	if err != nil {
		return nil, fmt.Errorf("读取证书密钥失败[%s]", r.KeyFile)
	}
	var ok bool
	switch r.KeyType {
	case FILE_CERT_CER:

		block, _ := pem.Decode(keyBuf)
		if len(block.Bytes) == 0 {
			return nil, fmt.Errorf("读取cer证书文件为空[%s]", r.KeyFile)
		}
		key, err := x509.ParseCertificate(block.Bytes) //解析cer
		if err != nil {
			return nil, fmt.Errorf("ParseCertificate err:%s", err)
		}
		r.PublicKey, ok = key.PublicKey.(*rsa.PublicKey)
		if !ok {
			return nil, fmt.Errorf("Value returned from ParsePKIXPublicKey was not an RSA public key")
		}
		r.Modules = (r.PublicKey.N.BitLen() + 7) / 8
	case FILE_CERT_PFX:
		prikey, _, err := pkcs12.Decode(keyBuf, cfg["KEY_PW"]) //第一个参数是*rsa.PrivateKey，第二个是CERTIFICATE类型的，里面也包含了公钥
		if err != nil {
			return nil, fmt.Errorf("解析PKCS12PFX失败[%s]", err)
		}
		r.PrivateKey, ok = prikey.(*rsa.PrivateKey)
		if !ok {
			return nil, fmt.Errorf("Value returned from ParsePrivateKey was not an RSA private key")
		}
		r.PublicKey = &(r.PrivateKey.PublicKey)
		r.Modules = (r.PrivateKey.N.BitLen() + 7) / 8
	default:
		return nil, fmt.Errorf("非法的RSA公钥类型[%s]", r.KeyType)
	}

	return r, nil
}

/*使用RSA公钥装载*/
func LoadFromRSAPublic(pub *rsa.PublicKey) (*RSAKey, error) {
	r := new(RSAKey)
	r.PublicKey = pub
	r.Modules = (r.PublicKey.N.BitLen() + 7) / 8
	return r, nil
}

/*使用RSA私钥装载*/
func LoadFromRSAPrivate(priv *rsa.PrivateKey) (*RSAKey, error) {
	r := new(RSAKey)
	r.PrivateKey = priv
	r.Modules = (r.PrivateKey.N.BitLen() + 7) / 8
	return r, nil
}

/*使用RSA公钥装载*/
func LoadRSAPubicPEM(buf string) (*RSAKey, error) {
	r := new(RSAKey)

	pemBuf, err := preloadPEM([]byte(buf))
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(pemBuf)
	if len(block.Bytes) == 0 {
		return nil, fmt.Errorf("读取RSA密钥文件w为空[%s]", r.KeyFile)
	}
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("解析RSA公钥失败[%s]", err)
	}
	var ok bool
	r.PublicKey, ok = pub.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("Value returned from ParsePKIXPublicKey was not an RSA public key")
	}
	r.Modules = (r.PublicKey.N.BitLen() + 7) / 8
	return r, nil
}

/*使用RSA私钥装载*/
func LoadRSAPrivatePEM(buf string) (*RSAKey, error) {
	r := new(RSAKey)
	var err error
	pemBuf, err := preloadPEM([]byte(buf))
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(pemBuf)
	if len(block.Bytes) == 0 {
		return nil, fmt.Errorf("读取RSA密钥文件w为空[%s]", r.KeyFile)
	}
	r.PrivateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("ParsePKCS1PrivateKey err:%s", err)
	}
	r.Modules = (r.PrivateKey.N.BitLen() + 7) / 8
	return r, nil
}

func (r *RSAKey) Sign(algo int, signbuf []byte) ([]byte, error) {
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
		hashType = crypto.MD5
		// return nil, x509.InsecureAlgorithmError(algo)
	default:
		return nil, fmt.Errorf("非法的签名算法:[%d]", algo)
	}

	if !hashType.Available() {
		return nil, x509.ErrUnsupportedAlgorithm
	}
	h := hashType.New()
	h.Write(signbuf)
	digest := h.Sum(nil)

	return r.PrivateKey.Sign(rand.Reader, digest, hashType)
}

func (r *RSAKey) Sha1(signbuf []byte) []byte {
	sha := sha1.New()
	sha.Write([]byte(signbuf))
	return []byte(hex.EncodeToString(sha.Sum(nil)))
}

func (r *RSAKey) Verify(algo int, signbuf []byte, signature []byte) error {
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
		hashType = crypto.MD5
	default:
		return fmt.Errorf("非法的验签算法:[%d]", algo)
	}

	if !hashType.Available() {
		return x509.ErrUnsupportedAlgorithm
	}
	h := hashType.New()
	h.Write(signbuf)
	digest := h.Sum(nil)

	return rsa.VerifyPKCS1v15(r.PublicKey, hashType, digest, signature)

}

func (r *RSAKey) EncryptPKCS1v15(origdata []byte) ([]byte, error) {
	encLen := r.Modules - 11 //PKCS1v15
	orig := bytes.NewBuffer(origdata)

	var ciperdata bytes.Buffer

	for {
		data := orig.Next(encLen)
		if len(data) == 0 {
			break
		}
		ciper, err := rsa.EncryptPKCS1v15(rand.Reader, r.PublicKey, data)
		if err != nil {
			return nil, fmt.Errorf("rsa.EncryptPKCS1v15 error[%s]", err)
		}
		ciperdata.Write(ciper)
	}
	return ciperdata.Bytes(), nil
}

func (r *RSAKey) EncryptPKCS1v15PRI(origdata []byte) ([]byte, error) {
	encLen := r.Modules - 11 //PKCS1v15
	orig := bytes.NewBuffer(origdata)

	var ciperdata bytes.Buffer

	for {
		data := orig.Next(encLen)
		if len(data) == 0 {
			break
		}
		ciper, err := rsa.SignPKCS1v15(rand.Reader, r.PrivateKey, crypto.Hash(0), []byte(data))
		if err != nil {
			return nil, fmt.Errorf("rsa.EncryptPKCS1v15 error[%s]", err)
		}
		ciperdata.Write(ciper)
	}
	return ciperdata.Bytes(), nil
}

func (r *RSAKey) DecryptPKCS1v15PUB(cipherdata []byte) ([]byte, error) {
	encLen := r.Modules //PKCS1v15
	orig := bytes.NewBuffer(cipherdata)

	var origdata bytes.Buffer

	for {
		data := orig.Next(encLen)
		if len(data) == 0 {
			break
		}
		ciper := rsa_public_decrypt(r.PublicKey, data)
		origdata.Write(ciper)
	}
	return origdata.Bytes(), nil

}

func rsa_public_decrypt(pubKey *rsa.PublicKey, data []byte) []byte {
	c := new(big.Int)
	m := new(big.Int)
	m.SetBytes(data)
	e := big.NewInt(int64(pubKey.E))
	c.Exp(m, e, pubKey.N)
	out := c.Bytes()
	skip := 0
	for i := 2; i < len(out); i++ {
		if i+1 >= len(out) {
			break
		}
		if out[i] == 0xff && out[i+1] == 0 {
			skip = i + 2
			break
		}
	}
	return out[skip:]
}

func (r *RSAKey) DecryptPKCS1v15(cipherdata []byte) ([]byte, error) {
	encLen := r.Modules //PKCS1v15
	orig := bytes.NewBuffer(cipherdata)

	var origdata bytes.Buffer

	for {
		data := orig.Next(encLen)
		if len(data) == 0 {
			break
		}
		ciper, err := rsa.DecryptPKCS1v15(rand.Reader, r.PrivateKey, data)
		if err != nil {
			return nil, fmt.Errorf("rsa.DecryptPKCS1v15 error[%s]", err)
		}
		origdata.Write(ciper)
	}
	return origdata.Bytes(), nil

}

func (r *RSAKey) GetModules() int {
	return r.Modules
}
