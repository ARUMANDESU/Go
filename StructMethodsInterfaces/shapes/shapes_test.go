package shapes

import "testing"

func TestArea(t *testing.T) {
	checkShapeArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("%#v got %.2f want %.2f", shape, got, want)
		}
	}

	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		checkShapeArea(t, rectangle, 72)
	})

	t.Run("circle", func(t *testing.T) {
		circle := Circle{10}
		checkShapeArea(t, circle, 314.1592653589793)
	})

	t.Run("table driven tests", func(t *testing.T) {
		areaTests := []struct {
			name  string
			shape Shape
			want  float64
		}{
			{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
			{name: "Circle", shape: Circle{Radius: 10}, want: 314.1592653589793},
			{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, want: 36.0},
		}

		for _, tt := range areaTests {
			t.Run(tt.name, func(t *testing.T) {
				checkShapeArea(t, tt.shape, tt.want)
			})

		}
	})
}
