func longestConsecutive(nums []int) int {
	m := map[int]int{}
	longest := 0
	for _, num := range nums {
		if _, exists := m[num]; exists {
			continue
		}

		length := m[num-1] + m[num+1] + 1
		m[num] = length
		m[num - m[num-1]] = length
		m[num + m[num+1]] = length

		longest = max(longest, length)
	}
	return longest
}
