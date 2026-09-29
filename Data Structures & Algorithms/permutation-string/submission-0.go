func checkInclusion(s1 string, s2 string) bool {
    freq := getFreq(s1)
    len1 := len(s1)
    len2 := len(s2)
    for i := 0; i < len2 - len1 + 1; i+= 1 {
        if getFreq(s2[i:i+len1]) == freq {
            return true
        }
    }
    return false
}

func getFreq(str string) [26]int {
    freq := [26]int{}
    for _, r := range str {
        freq[r - 'a'] += 1
    }
    return freq
}
