func trap(height []int) int {
    l, r := 0, len(height) - 1
    maxL, maxR := height[l], height[r]

    l += 1
    r -= 1
    water := 0
    for l <= r {
        if maxL <= maxR {
            water += max(0, maxL - height[l])
            maxL = max(maxL, height[l])
            l += 1
        } else {
            water += max(0, maxR - height[r])
            maxR = max(maxR, height[r])
            r -= 1
        }
    }

    return water
}
