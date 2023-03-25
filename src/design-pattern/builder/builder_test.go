package builder

import (
	"testing"
)


func TestNewPhoneByBuilder(t *testing.T) {
	tests := []struct {
		name        string
		builder IBuilder
	}{
		{
			name:        "HuaWei",
			builder: &HuaWeiBuilder{},
		},
		{
			name:        "Xiaomi",
			builder: &XiaomiBuilder{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			manager := Manager{builder: tt.builder}
			phone := manager.Produce()
			t.Logf("%v", phone)
		})
	}
}