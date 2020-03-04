package erratum

import (
	"fmt"
	"reflect"
)

func Use(opener func() (Resource, error), h string) error {
	var resource Resource
	var err error

	for {
		fmt.Println("Before opener()")
		resource, err = opener()
		fmt.Println("After opener(), before reflect")
		if reflect.TypeOf(err) != reflect.TypeOf(TransientError{}) {
			break
		}
	}


	resource.Frob(h)

	defer func() {
		fmt.Println("Before recover")
		if r := recover(); r != nil {
			if reflect.TypeOf(err) != reflect.TypeOf(FrobError{}) {
				resource.Defrob(h)
			}
		}
		fmt.Println("After recover")
	}()
	err = resource.Close()
	return err
}
