func checkInclusion(s1 string, s2 string) bool {
    len1 := len(s1)
    len2 := len(s2)
    if len2 < len1 {
        return false
    }
    freq := getFreq(s1)
    runningFreq := getFreq(s2[0:len1])
    matches := 0
    expected_matches := 0
    for i := range 26 {
        if freq[i] != 0 {
            expected_matches += 1
        }
    }
    for i := range 26 {
        if freq[i] != 0 && freq[i] == runningFreq[i] {
            matches += 1
        }
    }
    fmt.Println("Initial matches:", matches)
    fmt.Println("Expected matches:", expected_matches)
    for i := 0; i < len2 - len1 + 1; i+= 1 {
        fmt.Println("window:", s2[i:i+len1])
        fmt.Println("matches:", matches)
        if matches == expected_matches {
            return true
        }
        if i + len1 < len2 {
            a := s2[i]-'a'
            runningFreq[a] -= 1
            if freq[a] != 0 && freq[a] == runningFreq[a] {
                matches += 1
            } else if freq[a] != 0 && freq[a] == runningFreq[a] + 1 {
                matches -= 1
            }
            b := s2[i+len1]-'a'
            runningFreq[b] += 1
            if freq[b] != 0 && freq[b] == runningFreq[b] {
                matches += 1
            } else if freq[b] != 0 && freq[b] == runningFreq[b] - 1 {
                matches -= 1
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