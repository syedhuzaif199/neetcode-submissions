func groupAnagrams(strs []string) [][]string {
    m := map[[26]int][]int{}
    for i, str := range strs {
        freq := getFreq(str)
        if _, ok := m[freq]; !ok {
            m[freq] = []int{}
        }
        m[freq] = append(m[freq], i)
    }

    out := [][]string{}
    for key := range m {
        s := []string{}
        for _, i := range m[key] {
            s = append(s, strs[i])
        }
        out = append(out, s)
    }
    return out
}

func getFreq(s string) [26]int {
    out := [26]int{}
    for _, c := range s {
        out[c-'a'] += 1
    }
    return out
}