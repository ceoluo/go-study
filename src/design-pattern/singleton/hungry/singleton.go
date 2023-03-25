package hungry

import "fmt"

// 饿汉式 单例
type Singleton struct {

}

var singleton *Singleton

// init 函数是当所在的 package 首次被加载时执行，
// 若迟迟未被使用，则既浪费了内存，又延长了程序加载时间
func init()  {
	fmt.Printf("create a hungry singeleton instance \n")
	singleton = &Singleton{}
}

func GetSingletonInstance() *Singleton {
	return singleton
}
