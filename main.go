package main

import (
	"ascii-art/prepare"
	"fmt"
	"os"
	"strings"
)

const USAGE = `Usage: go run . [printed text] [color sequence/color name] [banner file name]\n\nEX: go run . "red blue" "111 6666"  "shadow"
colors:
0 -> grey, 1 -> red, 2 -> orange, 3 -> yellow, 4 -> green, 5 -> cyan, 6 -> blue, 7 -> purple, 8 -> white`

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
		bannerFileName = "standard.txt"
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

	switch lens {
	case 3, 4:
		fmt.Println(prepare.Color(words, lines, colorMask))
	default:
		fmt.Println(prepare.Aland(words, lines))

	}
	// for _, word := range words { // nested loop to print line by line depending on input.
	// 	if word == "" { // the new line "\\n" was at the end of "words" slice, and Split create the "" word
	// 		fmt.Println()
	// 	} else { // usual case letter print
	// 		// vertical step to print horizontal sequences of letter ascii art
	// 		for h := 1; h < 9; h++ { // from one to ignore the empty new line from standart.txt
	// 			for _, l := range word {
	// 				ind := (int(l)-32)*9 + h // potential index (the height from up to bottom) in "lines" for required letter line(because art letter is multilined)
	// 				if ind < len(lines) {    // check the index is inside available ascii art symbols ... f.e. standart.txt
	// 					fmt.Print(lines[ind]) // print the line from high "h" for the word letter "l"
	// 				}
	// 			}
	// 			fmt.Println()
	// 		}
	// 	}
	// }
}

func readFileIntoSlice(name string) []string {
	file, err := os.ReadFile(name)

	if err != nil {
		fmt.Printf("Error message: %s:\n", err)
		os.Exit(2)
	}
	lines := strings.Split(string(file), "\n")
	return lines
}
