package ascii

//ValidASCIIOutput - проверяем на валидность флажок
func ValidASCIIOutput(args []string) (bool, int) {
	for i, s := range args {
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
