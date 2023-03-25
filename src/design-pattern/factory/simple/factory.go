package simple

// 产品角色
type IRuleConfigParser interface {
	Parse(data []byte)
}


// 产品1
type jsonRuleConfigParser struct {

}

func (J *jsonRuleConfigParser) Parse(data []byte)  {
	panic("implement me")
}


// 产品2
type yamlRuleConfigParser struct {

}

func (Y *yamlRuleConfigParser) Parse(data []byte)  {
	panic("implement me")
}


// 新增产品3
//type xmlRuleConfigParser struct {
//
//}
//
//func (Y *xmlRuleConfigParser) Parse(data []byte)  {
//	panic("implement me")
//}


// 简单工厂
func NewIRuleConfigParser(t string) IRuleConfigParser {
	switch t {
	case "json":
		return &jsonRuleConfigParser{}
	case "yaml":
		return &yamlRuleConfigParser{}
	//case "xml":
	//	return &xmlRuleConfigParser{}
	}
	return nil
}
