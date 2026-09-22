func twoSum(numbers []int, target int) []int {
    p, q := 0, len(numbers) - 1

    for p < q {
        x := numbers[p] + numbers[q] - target
        if x == 0 {
            return []int{p+1, q+1}
        }
        if x > 0 {
            q -= 1
        } else {
            p += 1
        }
    }
    return []int{0, 0}
}