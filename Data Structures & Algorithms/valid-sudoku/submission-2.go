func isValidSudoku(board [][]byte) bool {
    // validate 3x3 squares
    for row := range 3 {
        for col := range 3 {
            occupied := [9]bool{}
            for i := range 3 {
                for j := range 3 {
                    num := board[3*row + i][3*col + j]
                    if num == '.' {
                        continue
                    }
                    num -= '1'
                    if occupied[num] {
                        return false
                    }
                    occupied[num] = true
                }
            }
        }
    }

    for row := range 9 {
        occupied := [9]bool{}
        for col := range 9 {
            num := board[row][col]
            if num == '.' {
                continue
            }
            num -= '1'

            if occupied[num] {
                return false
            }
            occupied[num] = true
        }
    }

    for col := range 9 {
        occupied := [9]bool{}
        for row := range 9 {
            num := board[row][col]
            if num == '.' {
                continue
            }
            num -= '1'

            if occupied[num] {
                return false
            }
            occupied[num] = true
        }
    }
    return true
}