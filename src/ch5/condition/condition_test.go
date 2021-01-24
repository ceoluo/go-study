package condition

import "testing"

func TestIfMutiSec(t *testing.T) {
	/*
		if condition{
			some operate
		}else if condition{
			some operate
		}else{
			some operate
		}

		if a:=condition; a{
			some operate
		}
	*/
	// 结合函数多返回值，判断执行是否成功
	if a := 1 == 1; a {
		t.Log("1==1")
	}

}

func TestSwitchMutiCase(t *testing.T) {
	for i := 0; i < 5; i++ {
		//switch i {
		switch {
		//case 0,2:
		case i%2 == 0:
			t.Log("Even")
		case i%2 == 1:
			t.Log("Odd")
		default:
			t.Log("It is not in 0-3")
		}
	}
}
