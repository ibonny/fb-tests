// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package Sample

import "strconv"

type Color int8

const (
	ColorRed   Color = 0
	ColorGreen Color = 1
	ColorBlue  Color = 2
)

var EnumNamesColor = map[Color]string{
	ColorRed:   "Red",
	ColorGreen: "Green",
	ColorBlue:  "Blue",
}

var EnumValuesColor = map[string]Color{
	"Red":   ColorRed,
	"Green": ColorGreen,
	"Blue":  ColorBlue,
}

func (v Color) String() string {
	if s, ok := EnumNamesColor[v]; ok {
		return s
	}
	return "Color(" + strconv.FormatInt(int64(v), 10) + ")"
}