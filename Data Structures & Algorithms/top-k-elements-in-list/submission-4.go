func topKFrequent(nums []int, k int) []int {
	freq := map[int]int{}

	for _, num := range nums {
		freq[num] += 1
	}

	bucketList := make([][]int, len(nums)+1)

	for key := range freq {
		bucketList[freq[key]] = append(bucketList[freq[key]], key)
	}

	ks := make([]int, 0, k)
	current := len(nums)
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
