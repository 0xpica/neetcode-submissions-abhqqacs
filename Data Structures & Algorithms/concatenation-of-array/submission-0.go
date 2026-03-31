func getConcatenation(nums []int) []int {
    l := len(nums)
    if l == 0 {
        return []int{}
    }
    ans := make([]int, l * 2)  
    for i := 0; i < l; i++ {
        ans[i]  = nums[i]
        ans[i+l] = nums[i]
    } 
    return ans
}
