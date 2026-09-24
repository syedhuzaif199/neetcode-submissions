func isValidSudoku(board [][]byte) bool {
    occRow := [9]int{}
    occCol := [9]int{}
    occSq  := [9]int{}

    for row := range 9 {
        for col := range 9 {
            num := board[row][col]
            if num == '.' {
                continue
            }
            num -= '1'

            mask := 1 << num
            if occCol[col] & mask != 0{
                return false
            }
            occCol[col] |= mask
            if occRow[row] & mask != 0 {
                return false
            }
            occRow[row] |= mask
            sqRow := row/3
            sqCol := col/3
            sqIdx := sqRow * 3 + sqCol
            if occSq[sqIdx] & mask != 0 {
                return false
            }
            occSq[sqIdx] |= mask
        }
    }
    return true
}
