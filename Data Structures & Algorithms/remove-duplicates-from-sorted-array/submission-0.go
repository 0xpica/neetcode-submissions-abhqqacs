func removeDuplicates(nums []int) int {
	x := 0
	for i := 1 ; i < len(nums) ; i++ {
		if nums[x] != nums[i] {
			x++
			nums[x] = nums[i]
		}
	}
	return x + 1 
}
