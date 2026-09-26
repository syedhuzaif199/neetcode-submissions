func reverse(x int) int {
	n := 0
	for x != 0 {
		if n > 214748364 || n < -214748364 {
            return 0
		}
		n = 10 * n + (x%10)
		x /= 10
	}
	return n
}
