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
		{
			"slices",
			[]Profile{
				{28, "Somewhere"},
				{33, "London"},
			},
			[]string{"Somewhere", "London"},
		},
		{
			"arrays",
			[2]Profile{
				{28, "Somewhere"},
				{33, "London"},
			},
			[]string{"Somewhere", "London"},
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

	// maps do not guarantee order
	// so they should just be tested if the value is present
	// in the []string
	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Cow":   "Moo",
			"Sheep": "Baa",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Moo")
		assertContains(t, got, "Baa")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{28, "Somewhere"}
			aChannel <- Profile{34, "Katowice"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Somewhere", "Katowice"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{28, "Somewhere"}, Profile{33, "London"}
		}

		var got []string
		want := []string{"Somewhere", "London"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false

	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}

	if !contains {
		t.Errorf("expected %v to contain %q but it didn't", haystack, needle)
	}
}
