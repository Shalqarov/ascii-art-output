package ascii_art

//ValidASCIIOutput - проверяем на валидность флажок
func ValidASCIIOutput(art ASCIIArt) (bool, int) {
	for i, s := range art.Argument {
		if len(s) > 9 {
			if s[:9] == "--output=" {
				if i != 0 {
					return true, i
				}
			}
		}
	}
	return false, 0
}
