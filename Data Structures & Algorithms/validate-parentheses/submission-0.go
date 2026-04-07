func isValid(strs string) bool {
	s := []rune{}
	for _ , c := range strs {
		if c == '[' || c == '{' || c == '(' { 
			s = append(s,c)
			continue
		}
		if len(s) == 0 {
			return false 
		}
		topIdx := len(s) -1 
		top := s[len(s)-1]
		if (c == ']' && top != '[') ||
		(c == '}' && top != '{') ||
		(c == ')' && top != '(') {
			return false
		} 
		s = s[:topIdx]

	}
	return len(s) == 0
}

