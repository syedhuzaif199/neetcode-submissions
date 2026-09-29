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
    longest := 0

    for q < len(s) {
        changes := q-p- maxi(freq)
        if changes <= k {
            longest = max(longest, q-p)
            freq[s[q]-'A'] += 1
            q += 1
        } else {
            freq[s[p] - 'A'] -= 1
            p += 1
        }
    }
    changes := q-p- maxi(freq)
    if changes <= k {
        longest = max(longest, q-p)
    }
    return longest
}