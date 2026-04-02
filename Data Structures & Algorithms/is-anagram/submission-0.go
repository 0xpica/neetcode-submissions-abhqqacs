func isAnagram(s string, t string) bool {
	sarr, tarr := [26]int{}, [26]int{}
	for _ , c := range s {
		sarr[c-'a']++
	}
	for _ , c := range t {
		tarr[c-'a']++
	}
	for i := 0 ; i < 26; i++ {
		if sarr[i] != tarr[i] {
			return false 
		}
	}
	return true
}
