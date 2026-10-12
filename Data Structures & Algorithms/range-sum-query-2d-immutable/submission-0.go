type NumMatrix struct {
    matrix [][]int
    summat [][]int
}

func Constructor(matrix [][]int) NumMatrix {
    out := NumMatrix{}
    out.matrix = matrix
    out.summat = make([][]int, len(matrix))

    for i, row := range matrix {
        sum := 0
        out.summat[i] = make([]int, len(row))
        for j, col := range row {
            sum += col
            out.summat[i][j] = sum
            if i > 0 {
                out.summat[i][j] += out.summat[i-1][j]
            }
        }
    }
    return out
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
    return this.sum(row2, col2) - this.sum(row1-1, col2) - this.sum(row2, col1 - 1) + this.sum(row1-1, col1-1)
}

// helper
func (this *NumMatrix) sum(row, col int) int {
    if row < 0 || col < 0 {
        return 0
    }
    return this.summat[row][col]
}

// Your NumMatrix object will be instantiated and called as such:
// obj := Constructor(matrix)
// param_1 := obj.SumRegion(row1,col1,row2,col2)
