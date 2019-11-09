package minsrv

type Mapping struct {
	m map[string]string
}

func NewMapping() *Mapping {
	return &Mapping{
		m: map[string]string{},
	}
}

func (m *Mapping) FileNameForPath(path string) string {
	return "/static/todo.html"
}

func (m *Mapping) HashedFileName(fileName string) string {
	return "/static/asgsfgasdfsd.html"
}
