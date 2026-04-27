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
	var key string
	var strArr []string
	for _ , s := range str {
		strArr = append(strArr,string(s))
	}
	sort.Strings(strArr)
	for _ , s := range strArr {
		key += s
	}
	return key
}
