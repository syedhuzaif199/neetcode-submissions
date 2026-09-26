func reverseBits(n int) int {
	res := 0
	left := 32
	for n > 0 {
		res = res << 1
		res |= n & 1
		n = n >> 1
		left -= 1
	}
	
	res = res << left
	return res
}
