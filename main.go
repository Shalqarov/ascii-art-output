package main

import (
	"fmt"
	"os"
	"strings"

	ascii "ascii/ascii_art"
)

func main() {
	art := ascii.ASCIIArt{Argument: os.Args[1:]}
	lenArgs := len(art.Argument)

	if lenArgs == 0 {
		fmt.Println()
		return
	}

	art.Output, _ = ascii.ValidASCIIOutput(art)
	_, idxOutput := ascii.ValidASCIIOutput(art)
	if !ascii.IsValidSym(art.Argument[0]) {
		fmt.Println("ascii-art: unregistered symbols")
		return
	}

	temp := strings.ReplaceAll(art.Argument[0], "\n", "\\n")
	art.Str = strings.Split(temp, "\\n")

	art.Fileinfo = ascii.OpenFont(art)

	if art.Output {
		ascii.Output(art, art.Argument[idxOutput][9:])
	} else {
		ascii.ConsoleOutput(art)
	}
}
