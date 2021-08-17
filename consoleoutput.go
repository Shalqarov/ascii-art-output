package ascii

import "fmt"

//ConsoleOutput - Вывод символов на консоль
func ConsoleOutput(art ASCIIArt) {
	for _, s := range art.Str { // цикл в аргументе, s - это строка в массиве
		if s != "" { // если строка не пустая, как раз на случай если передавать \n\n
			for i := 0; i < 8; i++ { // цикл на 8 строк
				for j := 0; j < len(s); j++ {
					fmt.Print(art.Fileinfo[int(s[j]-32)*9+1+i])
				}
				fmt.Println()
			}
		} else { //если передали \n\n
			fmt.Println()
		}
	}
}
