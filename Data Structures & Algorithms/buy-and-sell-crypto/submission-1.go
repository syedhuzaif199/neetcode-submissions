func maxProfit(prices []int) int {
    p, q := 0, 1
    max_profit := 0
    for q < len(prices) {
        profit := prices[q] - prices[p]
        if profit < 0 {
            p = q
        } else if profit > max_profit {
            max_profit = profit
        }
        
        q += 1
    }
    return max_profit
}
