package main
// 1、package名不一定和目录名相同，和Java不同
// 2、入口函数所在源码的package必须为main

import (
	"fmt"
	"os"
)

// main函数不支持传入参数，如Java, miain(string []args)
func main()  {
	// 在程序中直接通过os.Args来获取命令行参数，类比Python中的argparse
	fmt.Println("命令行参数以输出的格式获取：",os.Args)

	if len(os.Args)>1 {
		fmt.Printf("Hello %s \n", os.Args[1])
	}else {
		fmt.Println("Hello World!")
	}

	// main方法不能有返回值
	// 给出程序退出状态
	os.Exit(-1)
}
