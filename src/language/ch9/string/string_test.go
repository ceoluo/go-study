package string

import "testing"

func TestString(t *testing.T)  {
	// string是值类型，初始值为""
	var s string
	t.Log(s)

	s="hello"
	// s[1] = 3  // 出错，string是不可变的byte slice
	t.Log(len(s))

	s="\xE4\xB8\xA5" // 可以存储任何二进制数据

	t.Log(s)
	t.Log(len(s)) // 3 此处可证明字符串取len得到的不一定是字符的长度，是byte的长度

	s= "中"
	t.Log(len(s))

	c:=[]rune(s)
	t.Log(c)
	// t.log("rune size:", unsafeSizeof(c[0])
	t.Logf("中 Unicode %x", c[0])
	t.Logf("中 UTF8 %x", s)
}
