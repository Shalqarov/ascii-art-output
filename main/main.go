package main

import (
	"ascii"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	lenArgs := len(args)
	if lenArgs == 0 {
		fmt.Println()
		return
	}
	asciiOutput, idxOutput := ascii.ValidASCIIOutput(args)
	if !ascii.IsValidSym(args[0]) {
		fmt.Println("ascii-art: unregistered symbols")
		return
	}
	temp := strings.ReplaceAll(args[0], "\n", "\\n") //заменил все \n на \\n, чтобы была корректная строка, если не дописать кавычку
	argStr := strings.Split(temp, "\\n")             // разделил аргумент на строки
	str := []string{}
	str = ascii.OpenFont(args, asciiOutput)
	fmt.Print(asciiOutput)
	if asciiOutput {
		ascii.Output(argStr, str, args[idxOutput][9:])
	} else {
		ascii.ConsoleOutput(argStr, str)
	}

}
