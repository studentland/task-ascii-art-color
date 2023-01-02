# ascii-art-color
Colorized ASCII art generator. Background will be colored according incoming parameters.  

## Can not be used as task for curicculum, because not satisfying the requirements.
Created just for pleasure. As playing with available ways to colorize text in terminal.

![example.png][def]

To run the programm:  

- install `go` environment  
- clone the repo
- open terminal inside the repo root folder
- run `go run . "some text you want to print as art"` to print with Åland flag background and "standard.txt" font
- or `go run . "text" "" "shadow.txt"` to print with default background but with "shadow.txt" font
- or `go run . "text" "shadow.txt"` to print with Åland flag background and "shadow.txt" font
- or `go run . "one two" "123"` to print colored letters of word "one" as "o" red, "n" orange, "e" yellow, other letters as default

... more than three incoming params will be joined into one sequence, separated using space and printed as one word, with Åland flag background and "standard.txt" font.

Usage: `go run . [text] ?[colormask sequence or one color name] ?[font file name]`

Available colors -> color indices for background colormask sequence:
- default -> 0
- red -> 1
- orange -> 2
- yellow -> 3
- green -> 4
- cyan -> 5
- blue -> 6
- purple -> 7
- white -> 8
- dark -> 9

[def]: example.png