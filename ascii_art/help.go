package ascii_art

import (
	"fmt"
	"os"
)

//Help - --help
func Help(args []string) {
	for _, s := range args {
		if s == "--help" {
			fmt.Println("pomosh")
			os.Exit(0)
		}
	}
}
