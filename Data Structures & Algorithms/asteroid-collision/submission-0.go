func asteroidCollision(asteroids []int) []int {
    stack := make([]int, len(asteroids))
    top := 0

    for _, ast := range asteroids {
        stack[top] = ast
        top += 1
        for top > 1 && stack[top-1] < 0 && stack[top-2] > 0 {
            if stack[top-1] == -stack[top-2] {
                top -= 2
                continue
            }
            top -= 1
            if -stack[top] > stack[top-1] {
                stack[top-1] = stack[top]
            }
        }
    }

    return stack[:top]
}
