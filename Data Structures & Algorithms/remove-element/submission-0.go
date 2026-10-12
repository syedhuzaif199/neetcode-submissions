func removeElement(nums []int, val int) int {
	p, q := 0, 0

    for q < len(nums) {
        if nums[q] != val {
            nums[p] = nums[q]
            p += 1
        }

        q += 1
    }

    return p
 
}
