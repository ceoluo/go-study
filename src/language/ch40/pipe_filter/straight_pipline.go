package pipe_filter

type StraightPipline struct {
	Name    string
	Filters []Filter
}

func NewStraightPipline(name string, filters ...Filter) *StraightPipline {
	return &StraightPipline{Name: name, Filters: filters}
}

func (srf *StraightPipline) Process(data Request) (Response, error) {
	var ret interface{}
	var err error
	for _, filter := range srf.Filters {
		ret, err = filter.Process(data)
		if err != nil {
			return nil, err
		}
		data = ret
	}
	return ret, nil
}
