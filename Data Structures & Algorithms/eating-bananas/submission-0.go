func minEatingSpeed(piles []int, h int) int {
	largest := 0
	for _, num := range piles {
		largest = max(largest, num)
	}

	p, q := 1, largest
	k := largest

	for p <= q {
		rate := (p+q)/2
		available := h
		for _, num := range piles {
			available -= num/rate
			if num % rate != 0 {
				available -= 1
			}
		}

		if available >= 0 {
			k = min(k, rate)
			q = rate - 1
		} else {
			p = rate + 1
		}
	}

	return k
}
