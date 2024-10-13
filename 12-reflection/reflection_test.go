package main

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	// declaring and initializing an anonymous struct here
	// creates an anonymous struct with field Name of type string
	// and we are initializing it and assigning it to x by passing in expected, which is a string
	// x := struct {
	// 	Name string
	// }{expected}

	// since we are testing x which is an interface{}
	// it is better to use a table based test for different types

	// slice of structs where each struct has a Name, Input, and Expected Calls
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Number string
			}{"Samson"},
			[]string{"Samson"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Samson", "USA"},
			[]string{"Samson", "USA"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"Samson", 28},
			[]string{"Samson"},
		},
		{
			"nested fields",
			Person{
				"Samson",
				Profile{28, "Somewhere"},
			},
			[]string{"Samson", "Somewhere"},
		},
		{
			"pointers to things",
			&Person{
				"Samson",
				Profile{28, "Somewhere"},
			},
			[]string{"Samson", "Somewhere"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})
			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})

	}
}

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}
