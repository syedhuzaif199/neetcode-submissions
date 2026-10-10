type rect struct {
    w, h int
}
func largestRectangleArea(heights []int) int {
    stack := make([]rect, len(heights))
    top := 1
    stack[top-1] = rect{1, heights[0]}
    largest := 0

    for i := 1; i < len(heights); i+= 1 {
        h := heights[i]
        for top > 1 && h <= stack[top-1].h && h <= stack[top-2].h {
            top -= 1
            bigger := stack[top]
            smaller := stack[top-1]
            largest = max(largest, bigger.w * bigger.h)
            smaller.w += bigger.w
            stack[top-1] = smaller
        }

        last := stack[top-1]
        if h <= last.h {
            largest = max(largest, last.w * last.h)
            last.w += 1
            last.h = h
            stack[top-1] = last
        } else {
            stack[top] = rect{1, h}
            top += 1
        }
    }

    for top > 1 {
        top -= 1
        bigger := stack[top]
        smaller := stack[top-1]
        largest = max(largest, bigger.w * bigger.h)
        smaller.w += bigger.w
        stack[top-1] = smaller
    }

    last := stack[top-1]
    largest = max(largest, last.w * last.h)
    return largest
}