package print

import (
	"ascii-art/colors"
	"fmt"
	"math"
	"strings"
)

func PrepareAland(inputSlice []string, bannerRawLines []string) string {
	var bannerLines [][]rune
	for _, line := range bannerRawLines {
		bannerLines = append(bannerLines, []rune(line))
	}
	fmt.Println((bannerLines))
	var result string
	// if checker.AllEmpty(inputSlice) {
	// 	for l := 1; l < len(inputSlice); l++ {
	// 		result += "\n"
	// 	}
	// 	return result
	// }
	for i := 0; i < len(inputSlice); i++ {
		if len(inputSlice[i]) == 0 {
			result += "\n"
		} else {
			for j := 1; j < 9; j++ {
				for k := 0; k < len(inputSlice[i]); k++ {

					result += string(bannerLines[(int(inputSlice[i][k]-32)*9)+j])
				}
				result += "\n"
			}
		}
	}
	return colorResult(result)
}

const dxmin = 24 // min size for implementation something close to flag proportions
const dymin = 8  // will filled using spaces

func colorResult(s string) (news string) {
	splitted := strings.Split(s, "\n")
	dytext := len(splitted)
	var dxtext int // width of the area
	for _, sx := range splitted {
		xrunes := len([]rune(sx))
		if xrunes > dxtext { // search for the longest
			dxtext = xrunes
		}
	}

	dy := int(math.Max(float64(dytext), float64(dymin)))
	dx := int(math.Max(float64(dxtext), float64(dxmin)))

	colorStepX := dx / dxmin // how many columns includes one color step along x axis
	colorStepY := dy / dymin // how many rows includes one color step along y axis
	colorStep := int(math.Min(float64(colorStepX), float64(colorStepY)))

	si := symbolIndex(splitted, dx, dy)
	ci := colorIndex(colorStep, dx, dy)

	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			cIndex, _ := colors.IndexFromColorName(ci[y][x])
			color := colors.IColor(cIndex)
			news += color.Foreground + color.Background + si[y][x] + "\033[0m"
		}
		news += "\n"
	}
	return news
}

// symbolIndex calculates symbol for every cell of colored area
// [vertical row][horizontal column] symbol string
func symbolIndex(spl []string, dx, dy int) (si [][]string) {
	si = make([][]string, dy)
	for y := range si {
		si[y] = make([]string, dx)
	}
	lenY := len(spl)
	for y := 0; y < dy; y++ {
		var s string
		if y < lenY {
			s = spl[y]
		} else {
			s += " "
		}
		runes := []rune(s)
		lens := len(runes)
		for x := 0; x < dx; x++ {
			if x < lens {
				si[y][x] = string(runes[x])
			} else {
				si[y][x] = " "
			}
		}
	}
	return
}

// colorIndex calculates background color for every cell of colored area
// [vertical row][horizontal column] colorCodeString
func colorIndex(cs, dx, dy int) (ci [][]string) {
	ci = make([][]string, dy)
	for y := range ci {
		ci[y] = make([]string, dx)
	}

	// color map for zones
	cmap := [dymin][dxmin]string{
		{"blue", "blue", "blue", "blue", "yellow", "yellow", "red", "red", "red", "red", "yellow", "yellow", "blue", "blue", "blue", "blue", "blue", "blue", "blue", "blue", "blue", "blue", "blue", "blue"},
		{"blue", "blue", "blue", "blue", "yellow", "yellow", "red", "red", "red", "red", "yellow", "yellow", "blue", "blue", "blue", "blue", "blue", "blue", "blue", "blue", "blue", "blue", "blue", "blue"},

		{"yellow", "yellow", "yellow", "yellow", "yellow", "yellow", "red", "red", "red", "red", "yellow", "yellow", "yellow", "yellow", "yellow", "yellow", "yellow", "yellow", "yellow", "yellow", "yellow", "yellow", "yellow", "yellow"},
		{"red", "red", "red", "red", "red", "red", "red", "red", "red", "red", "red", "red", "red", "red", "red", "red", "red", "red", "red", "red", "red", "red", "red", "red"},

		{"red", "red", "red", "red", "red", "red", "red", "red", "red", "red", "red", "red", "red", "red", "red", "red", "red", "red", "red", "red", "red", "red", "red", "red"},
		{"yellow", "yellow", "yellow", "yellow", "yellow", "yellow", "red", "red", "red", "red", "yellow", "yellow", "yellow", "yellow", "yellow", "yellow", "yellow", "yellow", "yellow", "yellow", "yellow", "yellow", "yellow", "yellow"},
		{"blue", "blue", "blue", "blue", "yellow", "yellow", "red", "red", "red", "red", "yellow", "yellow", "blue", "blue", "blue", "blue", "blue", "blue", "blue", "blue", "blue", "blue", "blue", "blue"},
		{"blue", "blue", "blue", "blue", "yellow", "yellow", "red", "red", "red", "red", "yellow", "yellow", "blue", "blue", "blue", "blue", "blue", "blue", "blue", "blue", "blue", "blue", "blue", "blue"},
	}
	var stepX, stepY int
	for y := 0; y < dy; y++ {
		if (y+1)%cs == 0 && stepY < dymin-1 {
			stepY++
		}
		stepX = 0
		for x := 0; x < dx; x++ {
			if (x+1)%cs == 0 && stepX < dxmin-1 {
				stepX++
			}
			ci[y][x] = cmap[stepY][stepX]
		}
	}
	return
}
