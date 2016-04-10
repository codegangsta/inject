package inject_test

import (
	"fmt"
	"reflect"

	"github.com/codegangsta/inject"
)

func ExampleInjector() {
	// Create a new injector
	injector := inject.New()

	// Instantiate some dependency
	s := StructImplementingDependencyInterface{
		ID: 1,
	}
	s2 := StructImplementingDependencyInterface{
		ID: 2,
	}

	// Map your instantiated dependency to its interface type.
	// "(*DependencyInterface)(nil)" is a trick to get a pointer to an interface.
	injector.MapTo(s, (*DependencyInterface)(nil))
	// If you don't use interfaces, you can use Map instead of MapTo.
	injector.Map(s2)

	// Instantiate a struct needing dependencies, and do the injection.
	receiver := StructReceivingDependencies{}
	if err := injector.Apply(&receiver); err != nil {
		panic(err)
	}

	// As you can see in the output,
	// the interface field received the struct we mapped to the interface type,
	// while the struct field received the one we mapped to its actual type.
	fmt.Printf("interface field: %d\n", receiver.DependencyInterfaceField.GetID())
	fmt.Printf("struct field: %d\n", receiver.DependencyStructField.GetID())

	// Can also be used for functions.
	injector.Invoke(Print)

	// And finally, you can get the dependencies by using Get.
	structWeGot := injector.Get(reflect.TypeOf(StructImplementingDependencyInterface{})).Interface().(StructImplementingDependencyInterface)
	interfaceWeGot := injector.Get(inject.InterfaceOf((*DependencyInterface)(nil))).Interface().(DependencyInterface)

	fmt.Printf("interface we got: %d\n", interfaceWeGot.GetID())
	fmt.Printf("struct we got: %d\n", structWeGot.GetID())

	// output:
	// interface field: 1
	// struct field: 2
	// interface parameter: 1
	// struct parameter: 2
	// interface we got: 1
	// struct we got: 2
}

type DependencyInterface interface {
	GetID() int
}

type StructImplementingDependencyInterface struct {
	ID int
}

func (s StructImplementingDependencyInterface) GetID() int {
	return s.ID
}

type StructReceivingDependencies struct {
	DependencyInterfaceField DependencyInterface                   `inject:"-"`
	DependencyStructField    StructImplementingDependencyInterface `inject:"-"`
}

func Print(d1 DependencyInterface, d2 StructImplementingDependencyInterface) {
	fmt.Printf("interface parameter: %d\n", d1.GetID())
	fmt.Printf("struct parameter: %d\n", d2.GetID())
}
