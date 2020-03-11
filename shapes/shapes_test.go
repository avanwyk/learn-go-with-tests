package shapes

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	t.Run("should calculate perimeter of rectangle", func(t *testing.T) {
		actual := Perimeter(Rectangle{10.0, 20.0})
		expected := 60.0
		if expected != actual {
			t.Errorf("Expected %g, actual %g", expected, actual)
		}
	})
}

func TestArea(t *testing.T) {
	areaTestCases := []struct {
		name         string
		shape        Shape
		expectedArea float64
	}{
		{name: "Rectangle1", shape: Rectangle{20.0, 30.0}, expectedArea: 600.0},
		{name: "Rectangle2", shape: Rectangle{10.0, 20.0}, expectedArea: 200.0},
		{name: "Circle1", shape: Circle{20.0}, expectedArea: 1256.6370614359173},
		{name: "Circle2", shape: Circle{10.0}, expectedArea: math.Pi * 100},
		{name: "Triangle1", shape: Triangle{10.0, 10.0}, expectedArea: 50.0},
		{name: "Triangle2", shape: Triangle{20.0, 30.0}, expectedArea: 300.0},
	}

	for _, testArea := range areaTestCases {
		t.Run(testArea.name, func(t *testing.T) {
			shape := testArea.shape
			actual := shape.Area()
			if actual != testArea.expectedArea {
				t.Errorf("Shape %#v, Expected %g, actual %g", shape, testArea.expectedArea, actual)
			}
		})
	}
}
