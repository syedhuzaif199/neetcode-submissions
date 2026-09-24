func longestConsecutive(nums []int) int {
	longest := 0
	m := map[int]struct{}{}

	for _, num := range nums {
		m[num] = struct{}{}
	}

	for num := range m {
		if _, exists := m[num-1]; !exists {
			count := 1
			for {
				_, ok := m[num+1]
				if !ok {
					break
				}
				count += 1
				num += 1
			}
			longest = max(longest, count)
		}
	}

	return longest
}
