package security

import "encoding/base64"

func EncodeBase64(origdata []byte) []byte {
	str := base64.StdEncoding.EncodeToString(origdata)
	return []byte(str)
}

func DecodeBase64(cipherdata []byte) []byte {
	orig, err := base64.StdEncoding.DecodeString(string(cipherdata))
	if err != nil {
		return nil
	}
	return orig
}
