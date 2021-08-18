package ascii_art

//ValidASCIIOutput - function to find output banner, if it exists, returns true and index of argument
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

//IsValidSym - function for checking the validity of characters
func IsValidSym(s string) bool {
	for _, s := range s {
		if (s < 32 && s != 10) || s > 126 {
			return false
		}
	}
	return true
}
