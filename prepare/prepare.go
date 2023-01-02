package prepare

import (
	"ascii-art/colors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Color(words []string, lines []string, colorMask string) string {
	color, ok := colors.ColorNames[colorMask]
	if ok { // if colorMask is a color name
		return color + Aland(words, lines) + colors.ColorNames["reset"]
	} else { // if colorMask is a color indices sequence
		var output string
		colorIndices := strings.Split(colorMask, "")
		fmt.Println()
		for len(colorIndices) < len(strings.Join(words, "")) {
			colorIndices = append(colorIndices, "0") // append white color to the end of colorMask if it's length is less than the length of input
		}
		fmt.Println("colorIndices: ", colorIndices)
		var colorIndex int           // index of the current color in colorMask
		for _, word := range words { // nested loop to print line by line depending on input.
			if word == "" { // the new line "\\n" was at the end of "words" slice, and Split create the "" word
				output += "\n"
			} else { // usual case letter print
				var fresh = true // to print the message about unsupported symbols only once
				// vertical step to print horizontal sequences of letter ascii art
				for h := 1; h < 9; h++ { // from one to ignore the empty new line from standart.txt
					for _, l := range word {
						ind := (int(l)-32)*9 + h // potential index (the height from up to bottom) in "lines" for required letter line(because art letter is multilined)
						if ind < len(lines) {    // check the index is inside available ascii art symbols ... f.e. standart.txt
							// cIndex, _ := colors.OpositeColorNameIndexFromColorName(colors.)
							ci, _ := strconv.Atoi(colorIndices[colorIndex])
							name, _ := colors.ColorNameFromIndex(ci)
							oi, _ := colors.OpositeColorNameIndexFromColorName(name)
							color := colors.IColor(oi)
							output += color.Foreground + color.Background + lines[ind] + colors.Reset // print the line from high "h" for the word letter "l"
							colorIndex++
						} else {
							if fresh {
								fmt.Println("unsupported symbols was dropped")
								fresh = false
							}
						}
					}
					if h < 8 { // to print the next line of the same letter from the same color
						colorIndex -= len(word) // to print the next line of the same letter from the same color
					}
					output += "\n"
				}
			}
		}
		return output
	}
}

func Aland(inputSlice []string, bannerRawLines []string) string {
	var bannerLines [][]rune
	for _, line := range bannerRawLines {
		bannerLines = append(bannerLines, []rune(line))
	}
	var result string
	var fresh = true // to print the message about unsupported symbols only once
	for i := 0; i < len(inputSlice); i++ {
		if len(inputSlice[i]) == 0 {
			result += "\n"
		} else {
			for j := 1; j < 9; j++ {
				for k := 0; k < len(inputSlice[i]); k++ {
					ind := (int(inputSlice[i][k]-32) * 9) + j
					if ind < len(bannerLines) {
						result += string(bannerLines[ind])
					} else {
						result += "?"
						if fresh {
							fmt.Println("unsupported symbols was dropped")
							fresh = false
						}
					}
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
			cIndex, _ := colors.OpositeColorNameIndexFromColorName(ci[y][x])
			color := colors.IColor(cIndex)
			news += color.Foreground + color.Background + si[y][x] + colors.Reset
		}
		news += "\n"
	}
	return news
}

// symbolIndex calculates symbol for every cell of colored area
// [vertical row][horizontal column] symbol string
// if the text is shorter than the area, it will be filled with spaces
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

	// color map for zones of the flag of Aland Islands (Finland)
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
