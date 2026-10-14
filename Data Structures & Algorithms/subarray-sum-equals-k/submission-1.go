func subarraySum(nums []int, k int) int {
	m := map[int]int{}

	count := 0
	sum := 0
	for _, num := range nums {
		sum += num
		if sum == k {
			count += 1
		}
		if n, exists := m[sum-k]; exists {
			count += n
		}
		m[sum] += 1
	}

	return count
}
