package strings

func Reverse(s string) string {
	// This function takes a string and returns the reverse of it.
	// It does this by converting the string to a slice of runes
	// and then reversing
	slice := []rune(s)
	for i := 0; i < len(slice)/2; i++ {
		slice[i], slice[len(slice)-1-i] = slice[len(slice)-1-i], slice[i]
	}
	return string(slice)
}

func IsPalindrome(s string) bool {
	slice := []rune(s)
	for i := 0; i < len(slice)/2; i++ {
		if slice[i] != slice[len(slice)-i-1] {
			return false
		}
	}
	return true
}
