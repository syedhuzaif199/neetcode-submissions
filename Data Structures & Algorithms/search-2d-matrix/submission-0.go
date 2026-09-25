func searchMatrix(matrix [][]int, target int) bool {
    m, n := len(matrix), len(matrix[0])
    p, q := 0, m*n-1

    for p <= q {
        mid := (p+q)/2
        row := mid / n
        col := mid % n
        num := matrix[row][col]

        if num == target {
            return true
        }

        if num < target {
            p = mid + 1
        } else {
            q = mid - 1
        }
    }

    return false
}
