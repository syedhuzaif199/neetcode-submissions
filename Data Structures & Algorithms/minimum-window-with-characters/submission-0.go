func minWindow(s string, t string) string {
    freq := map[byte]int{}
    for i := range len(t) {
        freq[t[i]] += 1
    }

    expected := len(freq)
    running_freq := map[byte]int{}

    matches := 0

    var i, j int

    min_s := ""

    for j < len(s) {

        r := s[j]
        running_freq[r] += 1
        if _, exists := freq[r]; exists {
            if freq[r] == running_freq[r] {
                matches += 1
            }
        }
        j += 1
        
        for matches == expected {
            if min_s == "" || j-i < len(min_s) {
                min_s = s[i:j]
            }
            r := s[i]
            i += 1
            running_freq[r] -= 1
            if freq[r] == running_freq[r]+1 {
                matches -= 1
            }
        } 
    }

    return min_s
}

