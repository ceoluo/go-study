package abstractFactory

import (
	"testing"
)

// 用一个简单工厂封装工厂方法
func NewPhoneAndChargerFactory(t string) IFactory {
	switch t {
	case "HuaWei":
		return &HuaWeiPhoneFactory{}
	case "Xiaomi":
		return &XiaomiPhoneFactory{}
	}
	return nil
}

func TestNewPhoneAndChargerFactory(t *testing.T) {
	type args struct {
		t string
	}

	tests := []struct {
		name        string
		args        args
		phoneNumber string
	}{
		{
			name:        "HuaWei",
			args:        args{t: "HuaWei"},
			phoneNumber: "18208381239",
		},
		{
			name:        "Xiaomi",
			args:        args{t: "Xiaomi"},
			phoneNumber: "18208381239",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			phoneFactory := NewPhoneAndChargerFactory(tt.args.t)
			phone := phoneFactory.ProducePhone()
			phone.Call(tt.phoneNumber)
			phoneFactory.ProduceCharger().Charge(phone)
		})
	}
}
