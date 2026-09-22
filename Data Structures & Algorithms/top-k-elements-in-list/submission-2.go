func topKFrequent(nums []int, k int) []int {
	freq := map[int]int{}

	max := math.MinInt
	for _, num := range nums {
		freq[num] += 1
		if freq[num] > max {
			max = freq[num]
		}
	}

	bucketList := make([][]int, max+1)

	for key := range freq {
		bucketList[freq[key]] = append(bucketList[freq[key]], key)
	}

	ks := make([]int, 0, k)
	current := max
	count := k
	for count > 0 {
		for _, num := range bucketList[current] {
			if count <= 0 {
				break
			}
			ks = append(ks, num)
			count -= 1
		}
		current -= 1
	}

	return ks

}
