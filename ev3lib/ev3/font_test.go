package ev3

import "testing"

type FontListCoordTest struct {
	index int

	x int
	y int
}

func TestFontListToCoord(t *testing.T) {
	testData := make([]FontListCoordTest, 0)

	testData = append(testData, FontListCoordTest{0, 0, 0})
	testData = append(testData, FontListCoordTest{1, 1, 0})
	testData = append(testData, FontListCoordTest{16, 6, 1})

	for _, test := range testData {

		x, y := FontListToCoord(test.index)

		if x != test.x || y != test.y {
			t.Log("idx: ", test.index, " expected x: ", test.x, ", y: ", test.y, "\ngot x: ", x, ", y: ", y)
			t.Fail()
		}
	}
}
