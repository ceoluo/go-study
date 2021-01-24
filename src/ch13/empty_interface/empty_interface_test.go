package empty_interface

import (
	"fmt"
	"testing"
)

/*
go的接口的最佳实践：
	1、倾向于使用小的接口，很多接口只包含一个方法
		type Reader interface{
			Reader(p []byte) (n int, err error)
		}

		type Writer interface{
			Writer(p []byte) (n int, err error)
		}

	2、较大的接口定义，可以由多个小接口定义组合而成
		type ReadWriter interface{
			Reader
			Writer
		}

	3、只依赖于必要功能的最小接口
		func StoreData(reader Reader) error{
			...
		}
 */

// 空接口可以表示任何类型
func DoSomething(p interface{}) {
	// 使用断言来将接口转换为定制类型
	//if v, ok := p.(int); ok{
	//	fmt.Println("Integer", v)
	//	return
	//}
	//if v, ok := p.(string); ok {
	//	fmt.Println("string", v)
	//	return
	//}
	//fmt.Println("UnKnown Type")

	switch v := p.(type) {
	case int:
		fmt.Println("Integer", v)
	case string:
		fmt.Println("string", v)
	default:
		fmt.Println("UnKnown Type")
	}
}

func TestEmptyInterface(t *testing.T) {
	DoSomething(10)
	DoSomething("10")
}
