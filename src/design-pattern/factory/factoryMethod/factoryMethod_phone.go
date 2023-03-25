package factoryMethod

import "fmt"

// 抽象产品
type IPhone interface {
	Call(phoneNumber string)
}

// 抽象工厂
type IFactory interface {
	Produce() IPhone
}

// 具体产品1
// -----------------------------------
type HuaWeiPhone struct {
}

func (h *HuaWeiPhone) Call(phoneNumber string) {
	fmt.Println("华为手机拨打电话： " + phoneNumber)
}

// 具体工厂1
type HuaWeiPhoneFactory struct {
}

func (h *HuaWeiPhoneFactory) Produce() IPhone {
	return &HuaWeiPhone{}
}
// -----------------------------------


// 具体产品2
// -----------------------------------
type XiaomiPhone struct {
}

func (x *XiaomiPhone) Call(phoneNumber string) {
	fmt.Println("小米手机拨打电话： " + phoneNumber)
}

// 具体工厂2
type XiaomiFactory struct {
}

func (x *XiaomiFactory) Produce() IPhone {
	return &XiaomiPhone{}
}
// -----------------------------------


// 新增具体产品3
// -----------------------------------
type ApplePhone struct {
}

func (a *ApplePhone) Call(phoneNumber string) {
	fmt.Println("苹果手机拨打电话： " + phoneNumber)
}

type ApplePhoneParserFactory struct {
}

func (a *ApplePhoneParserFactory) Produce() IPhone {
	return &ApplePhone{}
}
// -----------------------------------

