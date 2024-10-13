package main

import (
	"reflect"
)

// we want to write a function walk(x interface{}, fn func(string)) which takes
// a struct `x` and calls `fn` for all string fields found inside
func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		switch field.Kind() {
		case reflect.String:
			fn(field.String())
		case reflect.Struct:
			walk(field.Interface(), fn)
		}
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	// can't use numField on pointer so need to extract it first
	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	return val
}
