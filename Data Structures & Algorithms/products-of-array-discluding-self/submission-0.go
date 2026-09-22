func productExceptSelf(nums []int) []int {
    pref := make([]int, len(nums))
    suff := make([]int, len(nums))
    prod := 1
    for i := range nums {
        pref[i] = prod
        prod *= nums[i]
    }

    prod = 1
    for i := range nums {
        idx := len(nums)-i-1
        suff[idx] = prod
        prod *= nums[idx]
    }

    out := make([]int, len(nums))
    for i := range out {
        out[i] = pref[i] * suff[i]
    }

    return out
}