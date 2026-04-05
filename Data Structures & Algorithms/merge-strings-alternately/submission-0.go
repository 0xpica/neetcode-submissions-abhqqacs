func mergeAlternately(word1 string, word2 string) string {
	ptr1 := 0
	ptr2 := 0 
	res := ""
	for ptr1 < len(word1) && ptr2 < len(word2) {
		if ptr1 <= ptr2 {
			res += string(word1[ptr1])
			ptr1++
		}else {
			res += string(word2[ptr2])
			ptr2++
		}
	}
	res += word1[ptr1:]
	res += word2[ptr2:]

	return res
}
