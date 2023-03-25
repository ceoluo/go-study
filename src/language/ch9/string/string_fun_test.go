package string

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringFun(t *testing.T)  {
	s := "A,B,C"
	// 字符串分割
	parts := strings.Split(s, ",")
	for _,part := range parts{
		t.Log(part)
	}
	t.Log(strings.Join(parts, "-"))
}

func TestStringConv(t *testing.T)  {
	// 整型转字符串
	s := strconv.Itoa(10)
	t.Log("sss"+s)

	// 字符串转整型
	s2 := "10"
	// strconv.Atoi(s2)返回两个值，需要判断成功与否
	if v,err := strconv.Atoi(s2);err==nil {
		t.Log(10+v)
	}
}
