func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for nIdx , n := range nums {
		rem := target - n
		if rIdx , ok := m[rem]; ok {
			return []int{rIdx,nIdx}
		}
		m[n] = nIdx
	}
	return []int{}
}
