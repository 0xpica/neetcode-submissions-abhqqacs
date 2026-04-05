func validPalindrome(s string) bool {
	if isPalindrome(s) {
		return true
	}
	for i := 0 ; i < len(s); i++ {
		tmp := s[0 : i ] + s[ i + 1 : ]
		if isPalindrome(tmp){
			return true
		}
	}
	return false
}

func isPalindrome(s string) bool {
	l , r := 0 , len(s) -1 
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}