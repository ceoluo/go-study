package extension

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Println("I'm Pet!")
}

func (p *Pet) SpeakTo(host string) {
	fmt.Print("Hello, ", host, ",")
	p.Speak()
}

type Dog struct {
	// 使用扩展的方式
	//p *Pet

	// 内嵌
	// 只能将Pet的方法嵌入到Dog中，不能重载，故无法将其当做继承，因为go不支持隐式类型转换，子类不能声明为父类类型（LSP）
	// 所以，想要使用方法重载，只能重新为新类型绑定方法
	//Pet
}

func (d *Dog) Speak() {
	// 扩展
	//d.p.Speak()
	// 重载
	fmt.Println("I'm Dog!")
}

func (d *Dog) SpeakTo(host string) {
	// 扩展
	//d.p.SpeakTo(host)
	fmt.Print("Hello, ", host, ",")
	d.Speak()
}

func TestDog(t *testing.T) {
	pet := new(Pet)
	pet.Speak()
	pet.SpeakTo("Dog")

	dog := new(Dog)
	dog.Speak()
	dog.SpeakTo("Pet")
}
