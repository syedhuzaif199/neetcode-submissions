func search(nums []int, target int) int {
	p, q := 0, len(nums)-1
	min_idx := 0
	for p <= q {
		mid := (p+q)/2
		if nums[mid] < nums[min_idx] {
			q = mid - 1
			min_idx = mid
		} else {
			p = mid + 1
		}
	}
	
	p, q = 0, len(nums)-1
	for p <= q {
		idx := (p+q)/2
		mid := (idx + min_idx) % len(nums)
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			p = idx + 1
		} else {
			q = idx - 1
		}
	}
	return -1
}
