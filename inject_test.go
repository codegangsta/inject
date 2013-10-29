package inject_test

import (
	"github.com/codegangsta/inject"
  "testing"
	"reflect"
)

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

  var result = ""
  err := injector.Invoke(func(dependency string) {
    result = dependency
  })

  expect(t, result, dep)
}
