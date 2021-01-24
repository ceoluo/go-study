package loop

import "testing"

func TestLoop(t *testing.T)  {
	n := 0
	// go没有while循环
	// 相当于 while(n<5){}
	for n<5 {
		t.Log(n)
		n++
	}
}
