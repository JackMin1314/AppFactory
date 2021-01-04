package security

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

type AESMOD int

const (
	AES_CBC_PKCS5PADDING AESMOD = iota
	AES_CBC_PKCS7PADDING
	AES_CBC_ZEROPADDING
	AES_ECB_PKCS5PADDING
	AES_ECB_PKCS7PADDING
	AES_ECB_ZEROPADDING
)

func AESEncryptIV(origData, key, iv []byte, mod AESMOD) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = getPadding(mod, origData, blockSize)
	blockModee := cipher.NewCBCEncrypter(block, iv[:blockSize])
	crypted := make([]byte, len(origData))
	blockModee.CryptBlocks(crypted, origData)
	return crypted, nil
}

func AESDecryptIV(crypted, key, iv []byte, mod AESMOD) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockModee := cipher.NewCBCDecrypter(block, iv[:blockSize])
	origData := make([]byte, len(crypted))
	blockModee.CryptBlocks(origData, crypted)
	return getUnpadding(mod, origData)
}

func AESEncrypt(origData, key []byte, mod AESMOD) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = getPadding(mod, origData, blockSize)
	var blockMode cipher.BlockMode
	switch mod {
	case AES_CBC_PKCS5PADDING, AES_CBC_PKCS7PADDING, AES_CBC_ZEROPADDING:
		blockMode = cipher.NewCBCEncrypter(block, key[:blockSize])
	case AES_ECB_PKCS5PADDING, AES_ECB_PKCS7PADDING, AES_ECB_ZEROPADDING:
		blockMode = NewECBEncrypter(block)
	}
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func AESDecrypt(crypted, key []byte, mod AESMOD) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	var blockMode cipher.BlockMode
	switch mod {
	case AES_CBC_PKCS5PADDING, AES_CBC_PKCS7PADDING, AES_CBC_ZEROPADDING:
		blockMode = cipher.NewCBCDecrypter(block, key[:blockSize])
	case AES_ECB_PKCS5PADDING, AES_ECB_PKCS7PADDING, AES_ECB_ZEROPADDING:
		blockMode = NewECBDecrypter(block)
	}
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	return getUnpadding(mod, origData)
}

func getPadding(mod AESMOD, data []byte, b_size int) []byte {
	switch mod {
	case AES_CBC_PKCS5PADDING, AES_ECB_PKCS5PADDING:
		return pPKCS5Padding(data, b_size)
	case AES_CBC_PKCS7PADDING, AES_ECB_PKCS7PADDING:
		return pPKCS7Padding(data, b_size)
	case AES_CBC_ZEROPADDING, AES_ECB_ZEROPADDING:
		return pZeroPadding(data, b_size)
	default:
		return nil
	}

}

func getUnpadding(mod AESMOD, data []byte) ([]byte, error) {
	switch mod {
	case AES_CBC_PKCS5PADDING, AES_ECB_PKCS5PADDING:
		return pPKCS5UnPadding(data)
	case AES_CBC_PKCS7PADDING, AES_ECB_PKCS7PADDING:
		return pPKCS7UnPadding(data)
	case AES_CBC_ZEROPADDING, AES_ECB_ZEROPADDING:
		return pZeroUnPadding(data)
	default:
		return nil, fmt.Errorf("invalid PADDING MOD[%d]", mod)
	}
}
