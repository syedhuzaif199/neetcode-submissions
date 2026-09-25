func search(nums []int, target int) int {
    p, q := 0, len(nums) - 1
    for p <= q {
        mid := (p+q)/2
        fmt.Println(p, q, mid, "-", nums[mid], target)
        if nums[mid] == target {
            return mid
        }

        if nums[mid] < target {
            p = mid + 1
        } else {
            q = mid - 1
        }
    }
    
    return -1

}
