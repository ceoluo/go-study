package encapsulation // 封装

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	ID string
	Name string
	Age int
}

// 为对象绑定方法

// 第一种方式，调用时实例会被拷贝一份
func (e Employee) String1() string  {
	fmt.Printf("String1 Adrees is %x\n", unsafe.Pointer(&e))
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.ID, e.Name, e.Age)
}

// 第二种方式，调用时使用原实例的指针，避免拷贝 (*推荐使用)
func (e *Employee) String2() string  {
	fmt.Printf("String2 Adrees is %x\n", unsafe.Pointer(e))
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.ID, e.Name, e.Age)
}

func TestCreateEmployeeObj(t *testing.T)  {
	// 实例创建及初始化
	e := Employee{"0", "Mike", 23}
	t.Log(e)
	t.Log(e.ID)
	t.Logf("e is %T", e)

	e1 := Employee{Name: "Jone", Age: 24}
	t.Log(e1)
	t.Log(e1.ID)
	t.Logf("e1 is %T", e1)

	// new 直接返回实例的指针，相当于 e2 := &Employee{}
	e2 := new(Employee)
	e2.ID = "2" // 与C++不同，通过实例的指针访问成员，不需要使用 ->
	e2.Name = "Rose"
	e2.Age = 25
	t.Log(e2)
	t.Logf("e2 is %T", e2)
}

func TestStructOperations(t *testing.T)  {
	e := Employee{"0", "Bob", 23}
	//e := &Employee{"0", "Bob", 23}
	fmt.Printf("Origin adrees is %x\n", unsafe.Pointer(&e))
	t.Log(e.String1())
	t.Log(e.String2())
}