package ev3

import (
	_ "embed"
	"math"
	"strings"
)

type RuneCoord struct {
	x, y int
}

//go:embed font.csv
var fontTable string

// Bool list is row first
var FontMap map[rune][]RuneCoord

const CharWidth int = 10
const CharHeight int = 16

func init() {
	FontMap = make(map[rune][]RuneCoord)

	lines := strings.Split(fontTable, "\n")

	for _, line := range lines {
		chars := strings.Split(line, ",")

		curr := make([]RuneCoord, 0)

		for i, a := range chars[1:] {
			if a == "1" {
				x, y := FontListToCoord(i)
				curr = append(curr, RuneCoord{x, y})
			}
		}
		var r rune
		for _, rr := range chars[0] {
			r = rr
			break
		}

		FontMap[r] = curr
	}
}

func FontListToCoord(index int) (x int, y int) {
	y = int(math.Floor((float64(index)) / float64(CharWidth)))
	x = int(math.Mod(float64(index), float64(CharWidth)))

	return
}
