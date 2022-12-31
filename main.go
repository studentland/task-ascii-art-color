package main

import (
	"ascii-art/print"
	"fmt"
	"os"
	"strings"
)

const USAGE = `Usage: go run . [printed text] [COLOR sequence/name] [banner name]\n\nEX: go run . "red blue" "111 6666"  "shadow"`

var unittest = false

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Please enter some text")
		return
	}
	input := strings.Join(os.Args[1:], " ")
	lines := readFileIntoSlice("standard.txt")
	runes := []rune(input) // convert input string to runes slice
	splittwo := "\\n"
	words := strings.Split(string(runes), splittwo)

	if input == "" {
		return
	}
	if input == "\\n" {
		fmt.Println()
		return
	}

	fmt.Println(print.PrepareAland(words, lines))

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
