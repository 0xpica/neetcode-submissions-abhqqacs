func subsetXORSum(nums []int) int {
	
	res := 0 

	var backtrack func(i int, sub []int)
	backtrack = func(i int, sub []int) {
		xor := 0 
		for _ , num := range sub {
			xor ^= num
		}
		res += xor 

		for j := i ; j < len(nums); j++ {
			sub = append(sub,nums[j])
			backtrack(j+1,sub)
			sub = sub[:len(sub)-1]
		}
	}

	backtrack(0,[]int{})

	return res

}
