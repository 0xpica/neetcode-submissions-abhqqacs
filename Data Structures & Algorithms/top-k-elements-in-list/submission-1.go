func topKFrequent(nums []int, k int) []int {
	// Step 1: frequency map
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}

	// Step 2: bucket (index = frequency)
	buckets := make([][]int, len(nums)+1)
	for num, f := range freq {
		buckets[f] = append(buckets[f], num)
	}

	// Step 3: collect top K
	res := make([]int, 0, k)
	for i := len(buckets) - 1; i >= 0 && len(res) < k; i-- {
		res = append(res, buckets[i]...)
	}

	return res[:k]
}