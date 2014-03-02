package inject

import (
	"fmt"
	"reflect"
	"strings"
)

type resolveError struct {
	chain   []reflect.Type
	fac     reflect.Value
	message string
}

func recoverResolvePanic(err *error) {
	if r := recover(); r != nil {
		switch x := r.(type) {
		case resolveError:
			chain := make([]string, len(x.chain))
			for i, t := range x.chain {
				chain[i] = fmt.Sprintf("%q", t)
			}

			if x.message == "" && x.fac.IsValid() {
				m := "factory 'func("
				facType := x.fac.Type()
				for i := 0; i < facType.NumIn(); {
					m += fmt.Sprintf("%v", facType.In(i))
					i += 1
				}
				m += ") " + fmt.Sprintf("%v", facType.Out(0)) + "'"
				x.message = m
			}

			*err = fmt.Errorf("Value not found for type %v (%v): %v",
				chain[len(chain) - 1], x.message, strings.Join(chain, " -> "))
			println(fmt.Sprintf("%v", *err))
		default:
			*err = fmt.Errorf("%v", x)
		}
	}
}
