func majorityElement(nums []int) int {
    ele := nums[0] 
	c := 1
	for i := 1; i < len(nums); i++ {
		if ele == nums[i] {
			c++
		}else {
			c-- 
			if c == 0 {
				ele = nums[i]
				c = 1 
			}
		}
	}
	return ele
}
