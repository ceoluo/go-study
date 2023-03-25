package abstractFactory

// 抽象产品1
type IRuleConfigParser interface {
	Parse(data []byte)
}


// 抽象产品2
type ISystemConfigParser interface {
	ParseSystem(data []byte)
}



// 抽象工厂
type IConfigParserFactory interface {
	CreateRuleConfigParser() IRuleConfigParser
	CreateSystemConfigParser() ISystemConfigParser
}


// 具体工厂1
type jsonConfigParserFactory struct {

}

// 具体产品A1
type jsonRuleConfigParser struct {

}

func (j *jsonRuleConfigParser) Parse(data []byte)  {
	panic("implement me")
}

// 具体产品A2
type jsonSystemConfigParser struct {

}

func (s *jsonSystemConfigParser) ParseSystem(data []byte)  {
	panic("implement me")
}

func (j *jsonConfigParserFactory) CreateRuleConfigParser() IRuleConfigParser {
	return &jsonRuleConfigParser{}
}

func (j *jsonConfigParserFactory) CreateSystemConfigParser() ISystemConfigParser {
	return &jsonSystemConfigParser{}
}


// 具体工厂2
type xmlConfigParserFactory struct {

}

// 具体产品B1
type xmlRuleConfigParser struct {

}

func (x *xmlRuleConfigParser) Parse(data []byte)  {
	panic("implement me")
}

// 具体产品B2
type xmlSystemConfigParser struct {

}

func (s *xmlSystemConfigParser) ParseSystem(data []byte)  {
	panic("implement me")
}

func (j *xmlConfigParserFactory) CreateRuleConfigParser() IRuleConfigParser {
	return &xmlRuleConfigParser{}
}

func (j *xmlConfigParserFactory) CreateSystemConfigParser() ISystemConfigParser {
	return &xmlSystemConfigParser{}
}

