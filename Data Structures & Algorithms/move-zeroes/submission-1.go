func moveZeroes(nums []int) {
	var i, j int

	for j < len(nums) {
		if nums[j] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i += 1
		}
		j += 1
	}
}
