package factoryMethod

// 抽象产品角色
type IRuleConfigParser interface {
	Parse(data []byte)
}

// 抽象工厂 （工厂方法接口）
type IRuleConfigParserFactory interface {
	CreateParser() IRuleConfigParser
}

// 具体产品1
type jsonRuleConfigParser struct {
}

func (J *jsonRuleConfigParser) Parse(data []byte) {
	panic("implement me")
}

// 具体工厂1
type jsonRuleConfigParserFactory struct {
}

func (J *jsonRuleConfigParserFactory) CreateParser() IRuleConfigParser {
	return &jsonRuleConfigParser{}
}

// 具体产品2
type yamlRuleConfigParser struct {
}

func (Y *yamlRuleConfigParser) Parse(data []byte) {
	panic("implement me")
}

// 具体工厂2
type yamlRuleConfigParserFactory struct {
}

func (Y *yamlRuleConfigParserFactory) CreateParser() IRuleConfigParser {
	return &yamlRuleConfigParser{}
}

// 新增具体产品3
type xmlRuleConfigParser struct {
}

func (Y *xmlRuleConfigParser) Parse(data []byte) {
	panic("implement me")
}

type xmlRuleConfigParserFactory struct {
}

func (X *xmlRuleConfigParserFactory) CreateParser() IRuleConfigParser {
	return &xmlRuleConfigParser{}
}

// 用一个简单工厂封装工厂方法
//func NewIRuleConfigParserFactory(t string) IRuleConfigParserFactory {
//	switch t {
//	case "json":
//		return &jsonRuleConfigParserFactory{}
//	case "yaml":
//		return &yamlRuleConfigParserFactory{}
//	case "xml":
//		return &xmlRuleConfigParserFactory{}
//	}
//	return nil
//}
