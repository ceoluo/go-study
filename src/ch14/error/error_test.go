package error

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

// go没有错误处理机制
// go通过errors.new("...")可产生一个错误
// go语言错误特性是快速失败
var LargerThanTwoError = errors.New("n should larger than 0!")
var LessThanOneHundredError = errors.New("n should less than 100!")

func GetFibonacci(n int) ([]int, error) {
	if n < 0 {
		return nil, LargerThanTwoError
	}
	if n > 100 {
		return nil, LessThanOneHundredError
	}
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

func GetFibonacci2(strn string) {
	// 尽早错误，避免嵌套
	var (
		i    int
		err  error
		list []int
	)
	if i, err = strconv.Atoi(strn); err != nil {
		fmt.Println("Error", err)
		return
	}
	if list, err = GetFibonacci(i); err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println(list)
}

func TestGetFibonacci(t *testing.T) {
	n := 10
	if fibList, err := GetFibonacci(n); err == LargerThanTwoError {
		t.Error("Less")
	} else if err == LessThanOneHundredError {
		t.Error("Larger")
	} else {
		t.Log(fibList)
	}

	GetFibonacci2("10")
}
