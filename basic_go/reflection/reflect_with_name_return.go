package reflection

import (
	"errors"
	"fmt"
	"reflect"
)

func reflect_test(a string, obj interface{}) error {
	rv := reflect.ValueOf(obj).Elem()
	a += "4"
	if !rv.CanSet() {
		return errors.New("can not set cache object value")
	}
	rv.Set(reflect.ValueOf(a))
	return nil
}

func load_cache(a string) (b string) {
	reflect_test(a, &b)
	return b
}

func Execute_reflect_with_name_return() {
	var a string = "123"
	c := load_cache(a)
	fmt.Println(a)
	fmt.Println(c)
}
