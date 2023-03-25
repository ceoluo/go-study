package constant

import (
	"testing"
)

/*
	常量定义
	* 与其它语言不同
		* 快速设置连续值
 */
const (
	Monday =  iota+1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

const (
	Readable = 1 << iota
	Writeable
	Executable
)

func TestConstant1(t *testing.T)  {
	t.Log(Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday)
}

func TestConstant2(t *testing.T)  {
	a := 7 // 0111
	t.Log(a&Readable == Readable, a&Writeable == Writeable, a&Executable == Executable)
}
