package _map

import "testing"

func TestMapWithFunValue(t *testing.T) {
	// map的value可以是函数类型
	// map实现工厂模式
	m := map[int]func(op int) int{}

	m[1] = func(op int) int {
		return op
	}

	m[2] = func(op int) int {
		return op * op
	}

	m[3] = func(op int) int {
		return op * op * op
	}
	t.Log(m[1](2), m[2](2), m[3](2))
}

func TestMapForSet(t *testing.T)  {
	// go中没有set类型，可以利用map实现一个set
	set := map[int]bool{}
	set[3] = true

	n := 3

	if set[3] {
		t.Logf("%d is existing ", n)
	}else {
		t.Logf("%d is not existing", n)
	}

	set[5] = true

	t.Log(len(set))

	delete(set, n)

	if set[3] {
		t.Logf("%d is existing ", n)
	}else {
		t.Logf("%d is not existing", n)
	}

	t.Log(set)
}