func isPalindrome(s string) bool {
	filter := ""
	for _ , c := range s {
		if ('a' <= c && c <= 'z') ||
			('A' <= c && c <= 'Z') {
				filter += strings.ToLower(string(c))
			}
		if ('0' <= c && c <= '9') {
			filter += string(c)
		}
	}
	l , r := 0 , len(filter) -1
	for l <= r {
		if filter[l] != filter[r] {
			return false
		}
		l++
		r--
	}
	return true 
}
