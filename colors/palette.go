package colors

import (
	"fmt"
)

type Color struct {
	Foreground string
	Background string
	Error      error
}

// IColor returns Color object, includes color codes given color index.
//
//	type Color struct {
//		Foreground string
//		Background string // oposite color for foreground
//		Error        error
//	}
//
// Indexes: 0: grey 1: red 2: orange 3: yellow 4: green 5: cyan 6: blue 7: purple 8: white
func IColor(i int) Color {
	foregroundColorName, err := colorNameFromIndex(i)
	if err != nil {
		return Color{Error: err}
	}
	foregroundColorCode, err := foregroundColor(foregroundColorName)
	if err != nil {
		return Color{Error: err}
	}
	backgroundColorName, err := opositeColorNameFromIndex(i)
	if err != nil {
		return Color{Error: err}
	}
	backgroundColorCode, err := backgroundColor(backgroundColorName)
	if err != nil {
		return Color{Error: err}
	}
	return Color{Foreground: foregroundColorCode, Background: backgroundColorCode, Error: err}
}

func opositeColorNameFromIndex(i int) (string, error) {
	ocn := map[int]string{
		0: "white",
		1: "green",
		2: "blue",
		3: "purple",
		4: "red",
		5: "orange",
		6: "orange",
		7: "yellow",
		8: "grey",
	}
	colorName, ok := ocn[i]
	if !ok {
		return "", fmt.Errorf("invalid oposite color index: %d", i)
	}
	return colorName, nil
}

func colorNameFromIndex(i int) (string, error) {
	cn := map[int]string{
		0: "grey",
		1: "red",
		2: "orange",
		3: "yellow",
		4: "green",
		5: "cyan",
		6: "blue",
		7: "purple",
		8: "white",
	}
	colorName, ok := cn[i]
	if !ok {
		return "", fmt.Errorf("invalid color index: %d", i)
	}
	return colorName, nil
}

func IndexFromColorName(s string) (int, error) {
	cn := map[string]int{
		"grey":   0,
		"red":    1,
		"orange": 2,
		"yellow": 3,
		"green":  4,
		"cyan":   5,
		"blue":   6,
		"purple": 7,
		"white":  8,
	}
	index, ok := cn[s]
	if !ok {
		return 0, fmt.Errorf("invalid color name: %s", s)
	}
	return index, nil
}

func OpositeColorNameIndexFromColorName(s string) (int, error) {
	i, err := IndexFromColorName(s)
	if err != nil {
		return i, err
	}
	oc, err := opositeColorNameFromIndex(i)
	if err != nil {
		return i, err
	}
	return IndexFromColorName(oc)
}

func foregroundColor(s string) (string, error) {
	cp := map[string]string{
		"grey":   "\033[30;1m",
		"red":    "\033[31;1m",
		"green":  "\033[32;1m",
		"yellow": "\033[33;1m",
		"blue":   "\033[34;1m",
		"purple": "\033[35;1m",
		"cyan":   "\033[36;1m",
		"white":  "\033[37;1m",
		"orange": "\033[38;2;255;165;0m", // todo test later with 1m
	}
	colorCode, ok := cp[s]
	if !ok {
		return "", fmt.Errorf("invalid foreground color name: %s", s)
	}
	return colorCode, nil
}

// Returns string if it is one of the valid options.
func backgroundColor(s string) (string, error) {
	cp := map[string]string{
		"grey":   "\033[40;1m",
		"red":    "\033[41;1m",
		"orange": "\033[48;2;255;165;0m", // todo test later with 1m
		"yellow": "\033[43;1m",
		"green":  "\033[42;1m",
		"cyan":   "\033[46;1m",
		"blue":   "\033[44;1m",
		"purple": "\033[45;1m",
		"white":  "\033[47;1m",
	}
	colorCode, ok := cp[s]
	if !ok {
		return "", fmt.Errorf("invalid background color name: %s", s)
	}
	return colorCode, nil
}

const Reset = "\033[0m"

var colorNames map[string]string = map[string]string{
	"grey":   "\033[30;1m",
	"red":    "\033[31;1m",
	"green":  "\033[32;1m",
	"yellow": "\033[33;1m",
	"blue":   "\033[34;1m",
	"purple": "\033[35;1m",
	"cyan":   "\033[36;1m",
	"white":  "\033[37;1m",
	"orange": "\033[38;2;255;165;0m", // todo test later with 1m
}
