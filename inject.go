package inject

import (
	"errors"
	"fmt"
	"reflect"
)

type Injector interface {
	Invoke(interface{}) error
	Apply(interface{}) error
	Map(interface{})
	MapTo(interface{}, interface{})
	Get(reflect.Type) reflect.Value
	SetParent(Injector)
}

type injector struct {
	values map[reflect.Type]reflect.Value
	parent Injector
}

func TypeOf(iface interface{}) reflect.Type {
	t := reflect.TypeOf(iface)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	return t
}

func New() Injector {
	return &injector{
		values: make(map[reflect.Type]reflect.Value),
	}
}

func (inj *injector) Invoke(f interface{}) error {
	t := reflect.TypeOf(f)

	var in = make([]reflect.Value, t.NumIn())
	for i := 0; i < t.NumIn(); i++ {
		argType := t.In(i)
		val := inj.Get(argType)
		if !val.IsValid() {
			return errors.New(fmt.Sprintf("Value not found for type %v", argType))
		}

		in[i] = val
	}

	reflect.ValueOf(f).Call(in)
	return nil
}

func (inj *injector) Apply(val interface{}) error {
	v := reflect.ValueOf(val)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return nil
	}

	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		structField := t.Field(i)
		if f.CanSet() && structField.Tag == "inject" {
			ft := f.Type()
			v := inj.Get(ft)
			if !v.IsValid() {
				return errors.New(fmt.Sprintf("Value not found for type %v", ft))
			}

			f.Set(v)
		}

	}

	return nil
}

func (i *injector) Map(val interface{}) {
	i.values[reflect.TypeOf(val)] = reflect.ValueOf(val)
}

func (i *injector) MapTo(val interface{}, ifacePtr interface{}) {
	i.values[TypeOf(ifacePtr)] = reflect.ValueOf(val)
}

func (i *injector) Get(t reflect.Type) reflect.Value {
	val := i.values[t]
	if !val.IsValid() && i.parent != nil {
		val = i.parent.Get(t)
	}
	return val
}

func (i *injector) SetParent(parent Injector) {
	i.parent = parent
}
