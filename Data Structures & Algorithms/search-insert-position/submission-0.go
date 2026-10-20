func searchInsert(nums []int, target int) int {
    p, q := 0, len(nums)-1
    
    for p <= q {
        mid := (p+q)/2
        if nums[mid] == target {
            return mid
        }
        if nums[mid] > target {
            q = mid - 1
        } else {
            p = mid + 1
        }
    }

    return p

}
