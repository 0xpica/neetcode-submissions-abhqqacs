func calPoints(op []string) int {
	stack := []int{}
	for _ , rec := range op {
		switch rec {
		case "+":
			top := len(stack) - 1
			add := stack[top] + stack[top-1]
			stack = append(stack, add)
		case "C":
			top := len(stack) - 1	
			stack = stack[:top]
		case "D":	
			top := len(stack) - 1	
			stack = append(stack,stack[top]*2)
		default:
			recInt, _ := strconv.Atoi(rec)
			stack = append(stack,recInt)
		}
	}
	res := 0 
	for i := 0 ; i < len(stack); i++ {
		res += stack[i]
	}
	return res
}
