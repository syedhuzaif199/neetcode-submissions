func rotate(nums []int, k int)  {
    k = k % len(nums)
    reverse(nums)
    reverse(nums[:k])
    reverse(nums[k:])
}

func reverse(arr []int) {
    p, q := 0, len(arr)-1
    for p < q {
        arr[p], arr[q] = arr[q], arr[p]
        p += 1
        q -= 1
    }
}
