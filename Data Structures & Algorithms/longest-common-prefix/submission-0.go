func longestCommonPrefix(strs []string) string {
	sort.Strings(strs)
	f := strs[0]
	l := strs[len(strs)-1]
	cl := min(len(f),len(l))
	res := ""
	for i := 0 ; i < cl ; i++ {
		if f[i] != l[i] {
			break
		}
		res += string(f[i])
	}
	return res
}
