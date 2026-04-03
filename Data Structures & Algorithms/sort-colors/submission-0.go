func sortColors(nums []int) {
    ic, jc, kc := 0 , 0 , 0
	for i := 0 ; i < len(nums); i++ {
		if nums[i] == 0 { ic++ }
		if nums[i] == 1 { jc++ }
		if nums[i] == 2 { kc++ }
	}
	for i := 0 ; i < ic ; i++ { nums[i] = 0 }
	for i := 0 + ic ; i < jc + ic ; i++ { nums[i] = 1 }
	for i := 0 + ic + jc ; i < kc + jc + ic ; i++ { nums[i] = 2 }
}
