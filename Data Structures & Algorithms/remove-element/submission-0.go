func removeElement(nums []int, val int) int {
	t := 0 
	for i := 0 ; i < len(nums); i++ {
		if val == nums[i] {
			continue
		}
		nums[t] = nums[i]
		t++
	}
	return t
}
