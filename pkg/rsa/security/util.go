package security

import (
	"crypto"
	"crypto/x509"
	"encoding/hex"
	"fmt"
)

func Hash(algo int, hashbuf []byte) ([]byte, error) {
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
	h.Write(hashbuf)
	digest := h.Sum(nil)

	return digest, nil
}

func EncodeHex(origdata []byte) []byte {
	str := hex.EncodeToString(origdata)
	return []byte(str)
}

func DecodeHex(hexdata []byte) []byte {
	data, err := hex.DecodeString(string(hexdata))
	if err != nil {
		return nil
	}
	return []byte(data)
}
