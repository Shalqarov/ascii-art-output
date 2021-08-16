package ascii

//isValidSym - проверяем, аргументы на невалидные символы
func IsValidSym(s string) bool {
	for _, s := range s { // Цикл на случай, если в агрументах будут символы, которых нет в нашем стандарте
		if (s < 32 && s != 10) || s > 126 {
			return false
		}
	}
	return true
}
