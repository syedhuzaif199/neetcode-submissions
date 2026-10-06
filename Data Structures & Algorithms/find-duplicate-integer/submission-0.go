func findDuplicate(nums []int) int {
    m := map[int]struct{}{}
	for _, num := range nums {
		if _, exists := m[num]; exists {
			return num
		}
		m[num] = struct{}{}
	}
	return 0
}
