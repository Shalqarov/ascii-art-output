package ascii_art

import (
	"bufio"
	"fmt"
	"os"
)

//OpenFont - opens file and write info from file line-by-line to var
func OpenFont(art ASCIIArt) []string {
	if len(art.Argument) == 1 || art.Output && len(art.Argument) == 2 {
		file, err := os.Open("fonts/standard.txt")
		if err != nil {
			fmt.Printf("ascii art fs: open '%s' no such file or directory\n", "standard")
			os.Exit(0)
		}
		defer file.Close()
		//scanning...
		scanner := bufio.NewScanner(file)
		str := []string{}
		//Writing file line-by-line to var str
		for scanner.Scan() {
			str = append(str, scanner.Text())
		}
		if len(str) != 855 {
			fmt.Println("ascii art: standard font is empty or it is not a font")
			os.Exit(0)
		}
		return str
	}

	file, err := os.Open("fonts/" + art.Argument[findFont(art)] + ".txt")
	if err != nil {
		fmt.Printf("ascii art fs: open '%s' no such file or directory\n", art.Argument[1])
		os.Exit(0)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	str := []string{}

	//Writing file line-by-line to var str
	for scanner.Scan() {
		str = append(str, scanner.Text())
	}

	if len(str) != 855 {
		fmt.Printf("ascii art: '%s' is empty or it is not a font\n", art.Argument[1])
		os.Exit(0)
	}

	return str
}

//findFont - returns index of font in our arguments
func findFont(art ASCIIArt) int {
	ishave, idx := ValidASCIIOutput(art)
	if ishave {
		for i := 1; i < len(art.Argument); i++ {
			if i == idx {
				continue
			} else {
				return i
			}
		}
	}
	return 1
}
