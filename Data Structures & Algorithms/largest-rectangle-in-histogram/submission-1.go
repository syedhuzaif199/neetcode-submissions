type pair struct {
    idx, h int
}

func largestRectangleArea(heights []int) int {
    stack := make([]pair, len(heights))
    top := 0
    largest := 0

    for i, h := range heights {
        idx := i
        for top > 0 && h < stack[top-1].h {
            last := stack[top-1]
            area := last.h * (i - last.idx)
            largest = max(largest, area)
            top -= 1
            idx = last.idx
        }

        stack[top] = pair{idx, h}
        top += 1
    }

    for top > 0 {
        last := stack[top-1]
        area := last.h * (len(heights) - last.idx)
        largest = max(largest, area)
        top -= 1
    }

    return largest

}