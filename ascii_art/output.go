package ascii_art

import (
	"fmt"
	"os"
)

// ASCIIArt - структура, сделал, чтобы легче читалось
type ASCIIArt struct {
	Argument []string
	Str      []string
	Fileinfo []string
	Output   bool
	Asciifs  bool
}

// Output ...
func Output(art ASCIIArt, filename string) {
	file, err := os.OpenFile(filename,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	for _, s := range art.Str {
		if s != "" {
			for i := 0; i < 8; i++ {
				for j := 0; j < len(s); j++ {
					file.WriteString(art.Fileinfo[int(s[j]-32)*9+1+i])
				}
				file.WriteString("\n")
			}
		} else {
			file.WriteString("\n")
		}
	}
}
