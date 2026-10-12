func sortColors(nums []int) {
    p := 0
    q := 0
    r := len(nums) - 1

    for q <= r {
        if nums[q] == 2 {
            nums[q] = nums[r]
            nums[r] = 2
            r -= 1
            continue
        }

        if nums[q] == 0 {
            nums[q] = 1
            nums[p] = 0
            p += 1
        }

        q += 1
    }
 
}
