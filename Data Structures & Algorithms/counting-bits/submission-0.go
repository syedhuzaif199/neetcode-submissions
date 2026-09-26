func countBits(n int) []int {
    out := make([]int, n+1)
    for i := 0; i <=n; i++ {
        x := i
        count := 0
        for x > 0 {
            if x & 1 == 1 {
                count += 1
            }
            x = x >> 1
        }
        out[i] = count
    }
    return out
}