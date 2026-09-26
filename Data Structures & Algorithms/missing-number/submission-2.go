func missingNumber(nums []int) int {
    out := len(nums)

    for i, num := range nums {
        out ^= i ^ num
    }
    return out
}
