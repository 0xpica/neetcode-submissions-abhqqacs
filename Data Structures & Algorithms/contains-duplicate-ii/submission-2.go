func containsNearbyDuplicate(nums []int, k int) bool {

	l := 0 
	window := map[int]bool{}
	for r := 0 ; r < len(nums); r++ {
		if r - l > k {
			delete(window,nums[l])
			l++
		}

		if window[nums[r]]{
			return true
		}
		window[nums[r]]=true
	}

	return false
}

