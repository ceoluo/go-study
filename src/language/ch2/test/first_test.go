package try_test

import "testing"

/*
注意：
	* 测试文件必须以_test.go结尾：xxx_test.go
	* 测试方法必须以Test开头,如: func TestXXX(t *testing.T) {...}
 */

func TestFirstTry(t *testing.T)  {
	t.Log("My first try")
}