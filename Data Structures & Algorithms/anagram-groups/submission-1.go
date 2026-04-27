func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _ , str := range strs {
		key := getKey(str)
		m[key] = append(m[key],str)
	}
	result := [][]string{}
	for _ , value := range m {
		result = append(result,value)
	}
	return result
}


func getKey(str string) string {
	characters := []rune(str)
	sort.Slice(characters, func(i, j int) bool {
		return characters[i] < characters[j]
	})
	return string(characters)
}
