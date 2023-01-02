package main

import (
	"ascii-art-color/colors"
	"ascii-art-color/prepare"
	"ascii-art-color/utils"
	"fmt"
	"os"
	"strings"
)

const USAGE = `Usage: go run . [printed text] [color sequence/color name] [banner file name]\n\nEX: go run . "red blue" "11106666" "shadow.txt"
colors:
0 -> default, 1 -> red, 2 -> orange, 3 -> yellow, 4 -> green, 5 -> cyan, 6 -> blue, 7 -> purple, 8 -> white, 9 -> dark`

var unittest = false

func main() {
	var input, colorMask, bannerFileName string
	var lines []string
	lens := len(os.Args)
	switch lens {
	case 1:
		fmt.Println(USAGE)
		return
	case 2:
		input = os.Args[1]
		colorMask = ""
		bannerFileName = "standard.txt"
	case 3:
		input = os.Args[1]
		colorMask = os.Args[2]
		if strings.HasSuffix(colorMask, ".txt") {
			bannerFileName = colorMask
			colorMask = ""
		} else {
			bannerFileName = "standard.txt"
		}
	case 4:
		input = os.Args[1]
		colorMask = os.Args[2]
		bannerFileName = os.Args[3]
	default:
		input = strings.Join(os.Args[1:], " ")
		colorMask = ""
		bannerFileName = "standard.txt"
	}
	lines = readFileIntoSlice(bannerFileName)
	splittwo := "\\n"
	words := strings.Split(input, splittwo)

	if input == "" {
		return
	}
	if input == "\\n" {
		fmt.Println()
		return
	}
	// todo: add color mask check
	if !utils.OnlyDigits(colorMask) {
		if _, ok := colors.ColorNames[colorMask]; !ok {
			fmt.Println(USAGE)
			return
		}
	}

	switch lens {
	case 3:
		if colorMask == "" {
			fmt.Println(prepare.Aland(words, lines))
		} else {
			fmt.Println(prepare.Color(words, lines, colorMask))
		}
	case 4:
		fmt.Println(prepare.Color(words, lines, colorMask))
	default:
		fmt.Println(prepare.Aland(words, lines))

	}
}

func readFileIntoSlice(name string) []string {
	file, err := os.ReadFile(name)
	if err != nil {
		fmt.Printf("Error message: %s:\n", err)
		os.Exit(2)
	}
	file = []byte(strings.ReplaceAll(string(file), "\r", ""))
	lines := strings.Split(string(file), "\n")
	return lines
}
