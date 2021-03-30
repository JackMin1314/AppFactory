package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {

	str1 := "hello"
	str2 := "world"
	fmt.Println(FMTSPrintf(str1, str2))
}

func FMTSPrintf(a, b string) string {
	return fmt.Sprintf("%s%s", a, b)
}

func AddOper(a, b string) string {
	return a + b
}

func StrJoin(a, b string) string {
	return strings.Join([]string{a, b}, "")
}

func ByteBuffer(a, b string) string {
	//定义Buffer类型
	var bt bytes.Buffer
	// 向bt中写入字符串
	bt.WriteString(a)
	bt.WriteString(b)
	//获得拼接后的字符串
	s3 := bt.String()
	return s3

}

func StringsBuilder(a, b string) string {
	var build strings.Builder
	build.WriteString(a)
	build.WriteString(b)
	s3 := build.String()
	return s3
}
