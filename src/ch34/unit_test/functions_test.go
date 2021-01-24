package testing

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	_ "github.com/stretchr/testify"
)

func TestSquare(t *testing.T)  {
	// 表格测试法
	inputs := [...]int{1,2,3}
	expected := [...]int{1,4,9}
	for i:= 0;i<len(inputs);i++ {
		ret := Square(inputs[i])
		if ret != expected[i] {
			t.Errorf("input is %d, the expected is %d, the actual %d",
				inputs[i], expected[i], ret)
		}
	}
}

func TestErrorInCode(t *testing.T)  {
	fmt.Println("Start")
	// 报错后继续执行
	t.Error("Error")
	fmt.Println("End")
}

func TestFatalInCode(t *testing.T)  {
	fmt.Println("Start")
	// 报错后停止执行
	t.Fatal("Error")
	fmt.Println("End")
}

func TestSquareWithAssert(t *testing.T)  {
	// 表格测试法
	inputs := [...]int{1,2,3}
	expected := [...]int{1,4,9}
	for i:= 0;i<len(inputs);i++ {
		ret := Square(inputs[i])
		assert.Equal(t, expected[i], ret)
	}
}

