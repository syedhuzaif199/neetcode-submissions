func checkInclusion(s1 string, s2 string) bool {
    len1 := len(s1)
    len2 := len(s2)
    if len2 < len1 {
        return false
    }
    freq := getFreq(s1)
    runningFreq := getFreq(s2[0:len1])
    for i := 0; i < len2 - len1 + 1; i+= 1 {
        if runningFreq == freq {
            return true
        }
        if i + len1 < len2 {
            runningFreq[s2[i]-'a'] -= 1
            runningFreq[s2[i+len1]-'a'] += 1
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
