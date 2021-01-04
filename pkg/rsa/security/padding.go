package security

import (
	"bytes"
	"fmt"
)

func pPKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pPKCS7UnPadding(plaintext []byte) ([]byte, error) {
	length := len(plaintext)
	unpadding := int(plaintext[length-1])
	rlen := length - unpadding
	if rlen <= 0 ||
		unpadding > 16 {
		return nil, fmt.Errorf("invalid PKCS7PADDING[%d][%02x]", unpadding, unpadding)
	}
	return plaintext[:rlen], nil
}

func pZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func pZeroUnPadding(plaintext []byte) ([]byte, error) {
	return bytes.TrimFunc(plaintext,
		func(r rune) bool {
			return r == rune(0)
		}), nil
}

func pPKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pPKCS5UnPadding(plaintext []byte) ([]byte, error) {
	length := len(plaintext)
	unpadding := int(plaintext[length-1])
	rlen := length - unpadding
	if rlen <= 0 ||
		unpadding > 16 {
		return nil, fmt.Errorf("invalid PKCS5PADDING[%d][%02x]", unpadding, unpadding)
	}
	return plaintext[:rlen], nil
}
