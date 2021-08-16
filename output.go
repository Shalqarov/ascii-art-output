package ascii

import (
	"fmt"
	"os"
)

func Output(str, fileinfo []string, filename string) {
	f, err := os.OpenFile(filename,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	for _, s := range str {
		if s != "" {
			for i := 0; i < 8; i++ {
				for j := 0; j < len(s); j++ {
					f.WriteString(fileinfo[int(s[j]-32)*9+1+i])
				}
				f.WriteString("\n")
			}
		} else {
			f.WriteString("\n")
		}
	}
}
