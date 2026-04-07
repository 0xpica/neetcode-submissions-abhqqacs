func merge(nums1 []int, m int, nums2 []int, n int) {
	temp := []int{}
	for i := 0 ; i < m; i++ {
		temp = append(temp,nums1[i])
	}
	for i := 0 ; i < n; i++ {
		temp = append(temp,nums2[i])
	}
	sort.Ints(temp)
	for i := 0 ; i < m + n ; i++ {
		nums1[i] = temp[i]
	}
}

