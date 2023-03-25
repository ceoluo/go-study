package series

import "fmt"

// 1 在main函数执行前，所有依赖的package的init方法都会被执行
// 2 导入的包的init只会会执行一次
// 3 一个包中可以定义多个init
func init()  {
	fmt.Println("init1")
}

func init()  {
	fmt.Println("init2")
}

func GetFibonacciSeries(n int) []int {
	ret := []int{1, 1}
	for i := 2; i < n; i++ {
		ret = append(ret, ret[i-2]+ret[i-1])
	}
	return ret
}

func Square(n int) int {
	return n * n
}
