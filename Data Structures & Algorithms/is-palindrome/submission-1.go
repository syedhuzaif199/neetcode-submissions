func isPalindrome(s string) bool {
    runes := []rune(s)
    p, q := 0, len(runes) -1

    for p < q {
        if !isAlphaNum(runes[p]) {
            p+=1
            continue
        }
        if !isAlphaNum(runes[q]) {
            q -= 1
            continue
        }

        if unicode.ToLower(runes[p]) != unicode.ToLower(runes[q]) {
            return false
        }
        p += 1
        q -= 1
    }
    return true
}

func isAlphaNum(r rune) bool {
    return unicode.IsLetter(r) || unicode.IsDigit(r)
}