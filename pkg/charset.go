package tools

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io"

	"github.com/axgle/mahonia"
)

var (
	GBKDecoder = mahonia.NewDecoder("GBK")
	GBKEncoder = mahonia.NewEncoder("GBK")
)

/*UTF8 转码到 GBK*/
func UTF8tGBK(src string) (dst []byte) {
	dst = []byte(GBKEncoder.ConvertString(src))
	return dst
}

/*GBK 转码到 UTF8*/
func GBKtUTF8(src []byte) (dst []byte) {
	tmp := GBKDecoder.ConvertString(string(src))
	return []byte(tmp)
}
func GBKReader(src string, r io.Reader) (io.Reader, error) {
	if GBKDecoder == nil {
		return nil, errors.New("非法编码字符集")
	}
	return GBKDecoder.NewReader(r), nil
}

// 返回XML GBK decoder; 调用 decoder.Decode()解析GBK编码的xml报文到结构体里
func XMLGBKDecoder(content []byte) *xml.Decoder {
	decoder := xml.NewDecoder(bytes.NewReader(content))
	// decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
	// 	return GBKReader(charset, input)
	// }
	decoder.CharsetReader = GBKReader
	return decoder
}

// 将utf-8的xml结构体字段用GBK编码并返回编码后的[]byte (not completely test)
func XMLGBKEncoder(v interface{}) []byte {
	buffer := &bytes.Buffer{}
	GBKEncoder.NewWriter(buffer)
	xmlEnc := xml.NewEncoder(buffer)
	// xmlEnc.Indent("", "	")
	if xmlEnc.Encode(v) != nil {
		return nil
	}
	return buffer.Bytes()
}

// 将string按照原格式解码并编码返回指定格式string (not completely test)
func ConvertToString(src string, oriCode string, desCode string) string {
	srcCoder := mahonia.NewDecoder(oriCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(desCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}
