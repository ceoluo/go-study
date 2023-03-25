package factoryMethod

import (
	"fmt"
	"reflect"
	"testing"
)

// 用一个简单工厂封装工厂方法
func NewIRuleConfigParserFactory(t string) IRuleConfigParserFactory {
	switch t {
	case "json":
		return &jsonRuleConfigParserFactory{}
	case "yaml":
		return &yamlRuleConfigParserFactory{}
	case "xml":
		return &xmlRuleConfigParserFactory{}
	}
	return nil
}

func TestNewIRuleConfigParserFactory(t *testing.T) {
	type args struct {
		t string
	}

	tests := []struct{
		name string
		args args
		want IRuleConfigParser
	}{
		{
			name: "json",
			args: args{t: "json"},
			want: &jsonRuleConfigParser{},
		},
		{
			name: "yaml",
			args: args{t: "yaml"},
			want: &yamlRuleConfigParser{},
		},
	}
	for _,tt := range tests{
		t.Run(tt.name, func(t *testing.T) {
			got := NewIRuleConfigParserFactory(tt.args.t).CreateParser()
			fmt.Printf("got: %v ， want： %v", got, tt.want)
			if got=got; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIRuleConfigParser() = %v, want %v", got, tt.want)
			}
		})
	}
}
