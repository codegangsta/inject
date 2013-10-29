package inject_test

import (
	"github.com/codegangsta/inject"
	"reflect"
	"testing"
)

type SpecialString interface {
}

/* Test Helpers */
func expect(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Errorf("Expected %v (type %v) - Got %v (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
}

func refute(t *testing.T, a interface{}, b interface{}) {
	if a == b {
		t.Errorf("Did not expect %v (type %v) - Got %v (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
}

func Test_InjectorInvoke(t *testing.T) {
	injector := inject.New()
	expect(t, injector == nil, false)

	dep := "some dependency"
	injector.Add(dep)
	dep2 := "another dep"
	injector.AddAs(dep2, (*SpecialString)(nil))

	err := injector.Invoke(func(d1 string, d2 SpecialString) {
		expect(t, d1, dep)
		expect(t, d2, dep2)
	})

	expect(t, err, nil)
}

func Test_TypeOf(t *testing.T) {
	iType := inject.TypeOf((*SpecialString)(nil))
	expect(t, iType.Kind(), reflect.Interface)

	iType = inject.TypeOf((*testing.T)(nil))
	expect(t, iType.Kind(), reflect.Struct)
}
