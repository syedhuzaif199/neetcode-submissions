func maxArea(height []int) int {
    p, q := 0, len(height) - 1
    maxWater := 0
    for p < q {
        water := (q - p) * min(height[q], height[p])
        maxWater = max(maxWater, water)

        if height[p] < height[q] {
            p += 1
        } else {
            q -= 1
        }
    }

    return maxWater
}