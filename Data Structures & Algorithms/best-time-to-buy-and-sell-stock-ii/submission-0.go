func maxProfit(prices []int) int {
    largest := 0
    smallest := prices[0]
    total := 0
    for i := 1; i < len(prices); i+=1 {
        profit := prices[i] - smallest
        if profit < largest {
            smallest = prices[i]
            total += largest
            largest = 0
            continue
        }
        largest = profit
    }

    total += largest 
    return total
}