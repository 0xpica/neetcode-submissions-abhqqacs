func containsNearbyDuplicate(nums []int, k int) bool {
	l := len(nums)
	for i := 0 ; i < l - 1 ; i++ {
		for j := i + 1; j < l; j++ {
			if nums[i] == nums[j] && abs(i , j) <= k {
				return true
			}
		}
	}
	return false
}

func abs(i,j int) int {
	if i > j {
		return i - j
	}
	return j - i
}