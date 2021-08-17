package ascii

import (
	"bufio"
	"fmt"
	"os"
)

//OpenFont - функция для открытия файла с шрифтом и чтения построчно
func OpenFont(art ASCIIArt) []string {
	if len(art.Argument) == 1 || art.Output && len(art.Argument) == 2 {
		file, err := os.Open("standard.txt") //открываем файл
		if err != nil {
			fmt.Printf("ascii art fs: open '%s' no such file or directory\n", "standard")
			os.Exit(0)
		}
		defer file.Close()                // отложенная функция закрытия файла
		scanner := bufio.NewScanner(file) //создаем переменную сканнер, сканируем файл
		str := []string{}
		cnt := 0 //переменная для подсчета строк
		for scanner.Scan() {
			str = append(str, scanner.Text()) // построчно записываем в массив строк . Каждая строка, как отдельный элемент массива
			cnt++
		}
		if cnt != 855 {
			fmt.Println("ascii art: standard font is empty")
			os.Exit(0)
		}
		return str
	}
	//fmt.Println(args[findFont(args)])
	file, err := os.Open(art.Argument[findFont(art)] + ".txt") //открываем файл
	if err != nil {
		fmt.Printf("ascii art fs: open '%s' no such file or directory\n", art.Argument[1])
		os.Exit(0)
	}
	defer file.Close()                // отложенная функция закрытия файла
	scanner := bufio.NewScanner(file) //создаем переменную сканнер, сканируем файл
	str := []string{}
	cnt := 0 //переменная для подсчета строк
	for scanner.Scan() {
		str = append(str, scanner.Text()) // построчно записываем в массив строк . Каждая строка, как отдельный элемент массива
		cnt++
	}
	if cnt != 855 {
		fmt.Printf("ascii art: '%s' is not a font\n", art.Argument[1])
		os.Exit(0)
	}

	return str
}

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
