func dailyTemperatures(temperatures []int) []int {
    out := make([]int, len(temperatures))

    stack := make([]int, len(temperatures))
    top := 0

    for i, temp := range temperatures {
        for top > 0 && temperatures[stack[top-1]] < temp {
            top -= 1
            out[stack[top]] = i - stack[top]
        }
        stack[top] = i
        top += 1
    }

    return out
}
