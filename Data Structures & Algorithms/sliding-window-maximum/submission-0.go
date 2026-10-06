func maxSlidingWindow(nums []int, k int) []int {
    out := make([]int, 0, len(nums) - k + 1)
    deque := make([]int, len(nums))
    dl := 0
    dr := 0

    p := 0
    q := 0
    for q < len(nums) {
        for dl < dr && nums[deque[dr-1]] < nums[q] {
            dr -= 1
        }
        if dl == dr {
            dl = 0
            dr = 0
        }

        deque[dr] = q
        dr += 1

        if p > deque[dl] {
            dl += 1
        }
        if dl == dr {
            dl = 0
            dr = 0
        }
        if (q+1) >= k {
            out = append(out, nums[deque[dl]])
            p += 1
        }

        q += 1
    }

    return out
}
