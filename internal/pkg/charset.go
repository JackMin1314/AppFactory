package tools

import (
	"errors"
	"io"

	"github.com/axgle/mahonia"
)

/*UTF8 转码到 GBK*/
func UTF8tGBK(src string) (dst []byte) {
	u2g := mahonia.NewEncoder("GBK")
	dst = []byte(u2g.ConvertString(src))
	return dst
}

/*GBK 转码到 UTF8*/
func GBKtUTF8(src []byte) (dst []byte) {
	u2g := mahonia.NewDecoder("GBK")
	tmp := u2g.ConvertString(string(src))
	return []byte(tmp)
}
func GBKReader(src string, r io.Reader) (io.Reader, error) {
	decoder := mahonia.NewDecoder("GBK")
	if decoder == nil {
		return nil, errors.New("非法编码字符集")
	}
	return decoder.NewReader(r), nil
}
