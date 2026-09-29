func maxi(nums [26]int) int {
    out := nums[0]
    for _, num := range nums {
        out = max(out, num)
    }
    return out
}
func characterReplacement(s string, k int) int {
    freq := [26]int{}

    var p, q int
    freq_sum := 0
    longest := 0

    for q < len(s) {
        changes := freq_sum - maxi(freq)
        if changes <= k {
            longest = max(longest, q-p)
            freq[s[q]-'A'] += 1
            freq_sum += 1
            q += 1
        } else {
            freq[s[p] - 'A'] -= 1
            freq_sum -= 1
            p += 1
        }
    }
    changes := freq_sum - maxi(freq)
    if changes <= k {
        longest = max(longest, q-p)
    }
    return longest
}