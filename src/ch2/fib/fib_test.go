package fib

import "testing"

func TestFibList(t *testing.T) {
	// 第一种变量声明和赋值方式
	//var a int = 1
	//var b int = 1

	// 第二种变量声明和赋值方式
	//var (
	//	a int = 1
	//	// go具有自动类型推断
	//	b = 1
	//)

	// 第三种变量声明和赋值方式
	//a := 1
	//b := 1

	// 在一个赋值语句中对多个变量进行赋值
	//var a,b int = 1,1
	//var a,b  = 1,1
	a, b := 1, 1

	t.Log(a)
	for i := 0; i < 5; i++ {
		t.Log(" ", b)
		a, b = b, a+b
	}
}
