package abstractFactory

import "fmt"

// 抽象产品1
type IPhone interface {
	Call(phoneNumber string)
	GetName() string
}

// 抽象产品2
type ICharger interface {
	Charge(phone IPhone)
}

// 抽象工厂
type IFactory interface {
	ProducePhone() IPhone
	ProduceCharger() ICharger
}

// -----------------------------------
// 具体产品1-1
type XiaomiPhone struct {
	name string
}

func (x *XiaomiPhone) Call(phoneNumber string) {
	fmt.Println("小米手机拨打电话： " + phoneNumber)
}

func (x *XiaomiPhone) GetName() string {
	return x.name
}

// 具体产品1-2
type XiaomiCharger struct {
}

func (x *XiaomiCharger) Charge(phone IPhone) {
	fmt.Printf("小米充电器给 %v 充电\n", phone.GetName())
}

// 具体工厂1
type XiaomiPhoneFactory struct {
}

func (x *XiaomiPhoneFactory) ProducePhone() IPhone {
	return &XiaomiPhone{name: "小米手机"}
}

func (x *XiaomiPhoneFactory) ProduceCharger() ICharger {
	return &XiaomiCharger{}
}
// -----------------------------------


// -----------------------------------
// 具体产品2-1
type HuaWeiPhone struct {
	name string
}

func (h *HuaWeiPhone) Call(phoneNumber string) {
	fmt.Println("华为手机拨打电话： " + phoneNumber)
}

func (h *HuaWeiPhone) GetName() string {
	return h.name
}

// 具体产品2-2
type HuaWeiCharger struct {
}

func (h *HuaWeiCharger) Charge(phone IPhone) {
	fmt.Printf("华为充电器给 %v 充电\n", phone.GetName())
}

// 具体工厂2
type HuaWeiPhoneFactory struct {
}

func (x *HuaWeiPhoneFactory) ProducePhone() IPhone {
	return &HuaWeiPhone{name: "华为手机"}
}

func (x *HuaWeiPhoneFactory) ProduceCharger() ICharger {
	return &HuaWeiCharger{}
}
// -----------------------------------