package inject

import (
	"errors"
	"reflect"
)

type Injector interface {
	Invoke(interface{}) error
	Add(interface{})
}

type injector struct {
	values map[reflect.Type]reflect.Value
}

func New() Injector {
	return &injector{
		values: make(map[reflect.Type]reflect.Value),
	}
}

func (inj *injector) Invoke(f interface{}) error {
  t := reflect.TypeOf(f)

  var in = make([]reflect.Value,t.NumIn())
  for i := 0; i < t.NumIn(); i++ {
    argType := t.In(i)
    val := inj.values[argType]
    if !val.IsValid() {
      return errors.New("TODO have a better error here")
    }

    in[i] = val
  }

  reflect.ValueOf(f).Call(in)
	return nil
}

func (i *injector) Add(d interface{}) {
	i.values[reflect.TypeOf(d)] = reflect.ValueOf(d)
}
