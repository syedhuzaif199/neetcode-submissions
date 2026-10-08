func characterReplacement(s string, k int) int {
    var p, q int

    freq := [26]int{}
    max_freq := 0
    longest := 0

    for q < len(s) {
        freq[s[q]-'A'] += 1
        max_freq = max(max_freq, freq[s[q]-'A'])
        
        if q-p+1-max_freq <= k {
            longest = max(longest, q-p+1)
        } else {
            freq[s[p]-'A'] -= 1
            p += 1
        }
        q += 1
    }
    return longest
}