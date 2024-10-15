package main

import (
	"reflect"
)

// we want to write a function walk(x interface{}, fn func(string)) which takes
// a struct `x` and calls `fn` for all string fields found inside
func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
		}
	case reflect.Chan:
		for {
			if v, ok := val.Recv(); ok {
				walkValue(v)
			} else {
				break
			}
		}
	case reflect.Func:
		valFnResult := val.Call(nil)

		for _, res := range valFnResult {
			walkValue(res)
		}
	}
}

func getValue(x interface{}) reflect.Value {
	// since x is basically an 'any' type
	// we have to see what the value of it is
	val := reflect.ValueOf(x)

	// can't use numField on pointer so need to extract it first
	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	return val
}
