package _type

import "testing"

/*
1、go数量类型：
	bool
	string
	int (默认int根据平台表现为int32或int64) int8 int16 int32 int64
	uint uint8 uint16 uint32 uint64 uintptr
	byte // 等同于uint8
	rune // 等同于int32, 代表一个Unicode编码的指针
	float32 float64
	complex64 complex128
	chan // chanel管道类型，协程通信csp时使用

2、与其它语言的主要差异：
	1、Go语言不支持隐式类型转换
	2、别名和原有类型也不能进行隐式类型转换

3、go不支持指针运算
4、go的string类型是值类型，也就是申明一个string变量时会被自动初始化为空字符串
*/

type MyInt int64

func TestImplicitType(t *testing.T) {
	var a int32 = 1
	var b int64
	// Go语言不支持隐式类型转换： cannot use a (type int32) as type int64 in assignment
	// b = a

	b = int64(a)

	var c MyInt
	// 别名和原有类型也不能进行隐式类型转换: cannot use b (type int64) as type MyInt in assignment
	//c = b

	c = MyInt(b)

	t.Log(a,b, c)
}
