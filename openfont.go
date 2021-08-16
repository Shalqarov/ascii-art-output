package ascii

import (
	"bufio"
	"fmt"
	"os"
)

//OpenFont - функция для открытия файла с шрифтом и чтения построчно
func OpenFont(args []string, haveBanner bool) []string {
	if len(args) == 1 || haveBanner && len(args) == 2 {
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
	file, err := os.Open(args[findFont(args)] + ".txt") //открываем файл
	if err != nil {
		fmt.Printf("ascii art fs: open '%s' no such file or directory\n", args[1])
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
		fmt.Printf("ascii art: '%s' is not a font\n", args[1])
		os.Exit(0)
	}

	return str
}

func findFont(args []string) int {
	ishave, idx := ValidASCIIOutput(args)
	if ishave {
		for i := 1; i < len(args); i++ {
			if i == idx+1 {
				continue
			} else {
				return i + 1
			}
		}
	}

	return 1
}
