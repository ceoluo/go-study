package polymorphsim

import (
	"fmt"
	"testing"
)

type Code string
// 接口，体现多态性
type Programmer interface {
	WriteHelloWorld() Code
}

type GoProgrammer struct {

}

func (goProgrammer *GoProgrammer)WriteHelloWorld() Code  {
	return "fmt.Println(\"Hello world!\")"
}

type JavaProgrammer struct {
	
}

func (javaProgrammer *JavaProgrammer) WriteHelloWorld() Code  {
	return "fmt.Println(\"Hello world!\")"
}

// 通过传入参数是接口类型和接口中定义同样的方法来对接口的实现
// 同时也可以看出空接口可表示任何类型
func WriteFirstProgrammer(p Programmer)  {
	fmt.Printf("%T %v\n", p, p.WriteHelloWorld())
}

func TestPolymorphism(t *testing.T)  {
	goProgrammer := new(GoProgrammer)
	javaProgrammer := new(JavaProgrammer)

	WriteFirstProgrammer(goProgrammer)
	WriteFirstProgrammer(javaProgrammer)
}