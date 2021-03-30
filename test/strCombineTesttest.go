package main

import (
	"testing"
)

//单元测试
func TestFMTSPrintf(t *testing.T) {
	//准备参数
	param1 := "hello"
	param2 := "world"
	//执行函数
	ret := FMTSPrintf(param1, param2)
	//判断结果是否符合预期
	if ret != "helloworld" {
		t.Error("FMTSPrintf result failed")
	}
}

func TestAddOper(t *testing.T) {
	//准备参数
	param1 := "hello"
	param2 := "world"
	//执行函数
	ret := AddOper(param1, param2)
	//判断结果是否符合预期
	if ret != "helloworld" {
		t.Error("AddOper result failed")
	}
}

func TestStrJoin(t *testing.T) {
	//准备参数
	param1 := "hello"
	param2 := "world"
	//执行函数
	ret := StrJoin(param1, param2)
	//判断结果是否符合预期
	if ret != "helloworld" {
		t.Error("StrJoin result failed")
	}
}

func TestByteBuffer(t *testing.T) {
	//准备参数
	param1 := "hello"
	param2 := "world"
	//执行函数
	ret := ByteBuffer(param1, param2)
	//判断结果是否符合预期
	if ret != "helloworld" {
		t.Error("ByteBuffer result failed")
	}
}

func TestStringsBuilder(t *testing.T) {
	//准备参数
	param1 := "hello"
	param2 := "world"
	//执行函数
	ret := StringsBuilder(param1, param2)
	//判断结果是否符合预期
	if ret != "helloworld" {
		t.Error("StringsBuilder result failed")
	}
}

func BenchmarkFMTSPrintf(b *testing.B) {
	//准备参数
	param1 := "hello"
	param2 := "world"
	//执行函数
	for i := 0; i < b.N; i++ {
		FMTSPrintf(param1, param2)
	}
}

//性能测试
func BenchmarkAddOper(b *testing.B) {
	//准备参数
	param1 := "hello"
	param2 := "world"
	//执行函数
	for i := 0; i < b.N; i++ {
		AddOper(param1, param2)
	}
}

func BenchmarkStrJoin(b *testing.B) {
	//准备参数
	param1 := "hello"
	param2 := "world"
	//执行函数
	for i := 0; i < b.N; i++ {
		StrJoin(param1, param2)
	}
}

func BenchmarkByteBuffer(b *testing.B) {
	//准备参数
	param1 := "hello"
	param2 := "world"
	//执行函数
	for i := 0; i < b.N; i++ {
		ByteBuffer(param1, param2)
	}
}

func BenchmarkStringsBuilder(b *testing.B) {
	//准备参数
	param1 := "hello"
	param2 := "world"
	//执行函数
	for i := 0; i < b.N; i++ {
		StringsBuilder(param1, param2)
	}
}
