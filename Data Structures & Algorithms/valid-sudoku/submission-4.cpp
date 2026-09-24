class Solution {
public:
    bool isValidSudoku(vector<vector<char>>& board) {
        int occRow[9] = {};
        int occCol[9] = {};
        int occSq[9]  = {};

        for (int row = 0; row < 9; row++) {
            for (int col = 0; col < 9; col++) {
                int num = board[row][col];
                if (num == '.') {
                    continue;
                }
                num -= '1';

                int mask = 1 << num;
                if ((occCol[col] & mask) != 0) {
                    return false;
                }
                occCol[col] |= mask;
                if ((occRow[row] & mask) != 0) {
                    return false;
                }
                occRow[row] |= mask;
                int sqRow = row/3;
                int sqCol = col/3;
                int sqIdx = sqRow * 3 + sqCol;
                if ((occSq[sqIdx] & mask) != 0) {
                    return false;
                }
                occSq[sqIdx] |= mask;
            }
        }
        return true;
    }
};
