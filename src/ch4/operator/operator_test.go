package operator

import "testing"

const (
	Readable = 1 << iota
	Writeable
	Executable
)

func TestCompareArray(t *testing.T) {
	/*
	go语言支持同维度，长度相同的数组比较
	如果维数和长度相同，且对应位置的元素等，则数组相等
	 */
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 4, 5}
	//c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}
	e := [...]int{3, 2, 1, 4}
	t.Log(a == b) // false
	//t.Log(a==c) // invalid operation: a == c (mismatched types [4]int and [5]int)
	t.Log(a == d) //true
	t.Log(a == e) //false
}

func TestConstant2(t *testing.T)  {
	a := 7 // 0111
	// 按位清零操作 &^ 异或
	t.Log(Readable, Writeable, Executable)
	t.Log(a&Readable==Readable, a&Writeable==Writeable, a&Executable==Executable)

	t.Log(a&^Readable)
	t.Log(a&^Readable == Readable) // 清除读权限

	t.Log(a&^Writeable == Writeable) // 清除写权限

	t.Log(a&^Executable == Executable) // 清除可执行权限
}