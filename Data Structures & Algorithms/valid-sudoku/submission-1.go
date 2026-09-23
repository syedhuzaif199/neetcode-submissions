func isValidSudoku(board [][]byte) bool {
    occRow := [9][9]bool{}
    occCol := [9][9]bool{}
    occSq  := [9][9]bool{}

    for row := range 9 {
        for col := range 9 {
            num := board[row][col]
            if num == '.' {
                continue
            }
            num -= '1'
            if occCol[col][num] {
                return false
            }
            occCol[col][num] = true
            if occRow[row][num] {
                return false
            }
            occRow[row][num] = true
            sqRow := row/3
            sqCol := col/3
            sqIdx := sqRow * 3 + sqCol
            if occSq[sqIdx][num] {
                return false
            }
            occSq[sqIdx][num] = true
        }
    }
    return true
}
