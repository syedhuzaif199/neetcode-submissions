func maxProfit(prices []int) int {
    pref := make([]int, len(prices))
    suff := make([]int, len(prices))

    // start with the smallest possible value to find the maximums
    x := 0
    for i, _ := range prices {
        idx := len(prices)-i-1
        suff[idx] = x
        x = max(x, prices[idx])
    }

    // x would be the max at this point, so can be used to find the minimums next
    for i, price := range prices {
        x = min(x, price)
        pref[i] = x
    }

    x = 0

    for i, _ := range prices {
        diff := suff[i] - pref[i]
        x = max(x, diff)
    }

    return x
}
