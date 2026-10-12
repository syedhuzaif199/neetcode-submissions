func majorityElement(nums []int) int {
	count := 0
    var majority int
    for _, num := range nums {
        if count == 0 {
            majority = num
            count += 1
            continue
        }
        if num == majority {
            count += 1
            continue
        }

        count -= 1
    }

    return majority

}
