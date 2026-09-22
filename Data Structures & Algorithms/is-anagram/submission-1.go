func isAnagram(s string, t string) bool {
    count := map[rune]int{}

    for _, i := range s {
        count[i] += 1
    }

    for _, i := range t {
        count[i] -=1
    }

    for key := range count {
        if count[key] != 0 {
            return false
        }
    }
    return true
}
