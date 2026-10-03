func findMin(nums []int) int {
    p, q := 0, len(nums)-1

    res := nums[0]
    for p <= q {
        mid := (p + q)/2
        if nums[mid] < res {
            q = mid - 1
            res = nums[mid]
        } else if nums[mid] >= res {
            p = mid + 1
        }
    }

    return res

}
