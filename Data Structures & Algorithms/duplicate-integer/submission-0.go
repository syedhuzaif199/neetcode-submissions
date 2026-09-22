import "slices"
func hasDuplicate(nums []int) bool {
    slices.Sort(nums)
    for i := 0; i < len(nums) - 1; i += 1{
        if nums[i] == nums[i+1] {
            return true
        }
    }
    return false
}
