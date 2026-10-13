func reverseString(s []byte)  {
    p := 0
    q := len(s) - 1
    for p < q {
        s[p], s[q] = s[q], s[p]
        p += 1
        q -= 1
    }
}
