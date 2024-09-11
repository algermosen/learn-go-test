package reflector

import "reflect"

func Walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)
	f := val.Field(0)
	fn(f.String())
}
