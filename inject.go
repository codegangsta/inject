package inject

type Injector interface {
	Invoke(interface{}) error
	Add(interface{})
}

type injector struct {
	values map[string]interface{}
}

func New() Injector {
	return &injector{}
}

func (i *injector) Invoke(f interface{}) error {
	return nil
}

func (i *injector) Add(d interface{}) {
}
