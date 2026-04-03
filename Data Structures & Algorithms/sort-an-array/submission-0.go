// selection sort: find min swap
func selectionSort(nums []int){
	l := len(nums)
	for i := 0; i < l-1; i++ {
		midx := i
		for j := i + 1; j < l; j++ {
			if nums[midx] > nums[j] {
				midx = j
			}
		}
		nums[i], nums[midx] = nums[midx], nums[i]
	}
}

// bubble sort: adjecent swap
func bubbleSort(nums []int) {
	l := len(nums)
	for i := 0; i < l; i++ {
		// to optmize, if alrady sorted.
		swapped := false
		for j := 0; j < l-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

// merge sort: devide and sort
func mergeSort(nums []int) []int{
	if len(nums) <= 1 {
		return nums
	}
	mid := len(nums) / 2 

	left := mergeSort(nums[:mid])
	right := mergeSort(nums[mid:])
	
	return merge(left,right)
}

func merge(left, right []int) []int{
	result := make([]int,0, len(left)+len(right))
	i , j := 0 , 0 
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++ 
		}else{
			result = append(result, right[j])
			j++
		}
	}
	result = append(result,left[i:]...)
	result = append(result, right[j:]...)
	return result
}

func sortArray(nums []int) []int {
	return mergeSort(nums)
}
