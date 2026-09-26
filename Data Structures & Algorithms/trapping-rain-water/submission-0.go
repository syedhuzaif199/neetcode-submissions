func trap(height []int) int {
    greatestToLeft  := make([]int, len(height))
    greatestToRight := make([]int, len(height))
    greatest := 0

    for i, h := range height {
        greatestToLeft[i] = greatest
        greatest = max(greatest, h)
    }

    greatest = 0
    for i, _ := range height {
        idx := len(height) - i - 1
        greatestToRight[idx] = greatest
        greatest = max(greatest, height[idx])
    }

    water := 0
    for i, h := range height {
        water += max(0, min(greatestToLeft[i], greatestToRight[i]) - h)
    }
    return water
}

