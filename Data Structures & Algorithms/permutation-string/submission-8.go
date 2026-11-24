func checkInclusion(s1 string, s2 string) bool {
    len1 := len(s1)
    len2 := len(s2)
    if len2 < len1 {
        return false
    }
    freq := getFreq(s1)
    runningFreq := getFreq(s2[0:len1])
    diff := 0
    for i := range 26 {
        if freq[i] != runningFreq[i] {
			diff += 1
        }
    }

    for i := 0; i < len2 - len1 + 1; i+= 1 {
        if diff == 0 {
            return true
        }
        if i + len1 < len2 {
            a := s2[i]-'a'
            runningFreq[a] -= 1
            if freq[a] == runningFreq[a] {
                diff -= 1
            } else if freq[a] == runningFreq[a] + 1 {
                diff += 1
            }
            b := s2[i+len1]-'a'
            runningFreq[b] += 1
            if freq[b] == runningFreq[b] {
                diff -= 1
            } else if freq[b] == runningFreq[b] - 1 {
                diff += 1
            }
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