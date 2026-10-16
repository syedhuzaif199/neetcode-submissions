func getConcatenation(nums []int) []int {
    out := make([]int, 2*len(nums))
	for i, num := range nums {
		out[i] = num
		out[i+len(nums)] = num
	}

	return out
}
