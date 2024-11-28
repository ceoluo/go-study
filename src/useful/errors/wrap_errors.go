package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
)

// 这里存在一个重复wrap的问题，后面会说明如何避免。

// 这里只会简单的说明几个关键方法的用法，后文会分析它们的代码。

// 整个调用链是 main -> func3 -> func2 -> func1 main是调用链源头
// func1是出现根错误的地方，func2调用func1出现错误后秉承只处理一次的原则，直接return。func3类似操作。
// main作为调用链的源头，处理error的方式选择的是打印日志。
func main() {
	err := func3()
	if err != nil {
		// Cause可以找到被Wrap的根err
		log.Printf("original error:%T %v\n", errors.Cause(err), errors.Cause(err))
		// 打印Wrap过的err 可以通过%+v获取到调用的堆栈信息
		log.Printf("stack trace:\n%+v\n", err)
		os.Exit(1)
	}
}

func func3() error {
	err := func2()
	if err != nil {
		// Wrap error可以携带上下文
		return errors.Wrap(err, "fun3 call func2 error")
	}
	fmt.Println("i am func3")
	return nil
}

func func2() error {
	err := func1()
	if err != nil {
		// Wrap error可以携带上下文
		return errors.Wrap(err, "func2 call func1 error")
	}
	fmt.Println("i am func2")
	return nil
}

func func1() error {
	// 产生根error
	return errors.New("fun1 root error")
}
