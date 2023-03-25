package builder

import "fmt"

type Phone struct {
	// 名称
	name string
	// 处理器
	processor string
	// 摄像头
	camera string
	// 屏幕
	screen string
}


type IBuilder interface {
	AddProcessor(phone *Phone)
	AddCamera(phone *Phone)
	AddScreen(phone *Phone)
	Produce() *Phone
}

// ------------------------------------------------
type HuaWeiBuilder struct {

}

func (h *HuaWeiBuilder) AddProcessor(phone *Phone) {
	phone.processor = "海思麒麟处理器"
}

func (h *HuaWeiBuilder) AddCamera(phone *Phone) {
	phone.camera = "莱卡摄像头"
}

func (h *HuaWeiBuilder) AddScreen(phone *Phone) {
	phone.screen = "OLED"
}

func (h *HuaWeiBuilder) Produce() *Phone {
	phone := &Phone{name: "HuaWeiPhone"}
	h.AddProcessor(phone)
	h.AddCamera(phone)
	h.AddScreen(phone)
	fmt.Printf("华为生产线的工人组装了一台华为手机\n")
	return phone
}
// ------------------------------------------------


// ------------------------------------------------
type XiaomiBuilder struct {

}

func (x *XiaomiBuilder) AddProcessor(phone *Phone) {
	phone.processor = "高通骁龙处理器"
}

func (x *XiaomiBuilder) AddCamera(phone *Phone) {
	phone.camera = "索尼摄像头"
}

func (x *XiaomiBuilder) AddScreen(phone *Phone) {
	phone.screen = "OLED"
}

func (x *XiaomiBuilder) Produce() *Phone {
	phone := &Phone{name: "XiaomiPhone"}
	x.AddProcessor(phone)
	x.AddCamera(phone)
	x.AddScreen(phone)
	fmt.Printf("小米生产线的工人组装了一台小米手机\n")
	return phone
}
// ------------------------------------------------


type Manager struct {
	builder IBuilder
}

func (m *Manager) Produce() *Phone {
	return m.builder.Produce()
}
