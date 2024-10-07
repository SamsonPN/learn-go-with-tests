package main

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		// 2 decimal places for the float64 type
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	// we created this function to reduce duplication of our code
	// all we had to do was call this and pass in Shape type
	// to find its area
	// we do not have to call rectangle.Area() or circle.Area() directly

	// if we try to pass in something that isn't a Shape
	// then it would not compile

	// we have essentially decoupled this helper from concrete types
	// and only care about the method we are testing!!!
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		checkArea(t, rectangle, 72.0)
	})
	t.Run("circle", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})

	// instead of these 2 tests that essentially do the same thing
	// we can create something called a Table Driven Test

	// slice of anonymous structs
	// using this pattern, we can easily add a new Shape
	// and test its area method without having to create a new test
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		// naming the fields for creating a new struct
		// makes things clear what each value means
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {

		// for Table Driven Tests, use t.Run to name
		// the test cases which makes it clear
		// which test failed
		// use tt.name as the t.Run test name

		// can also be used to run specific tests
		// within the table with
		// go test -run TestArea/[test-name]
		// e.g. go test -run TestArea/Rectangle
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				// %#v prints out the struct with the values in its field
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})
	}
}
