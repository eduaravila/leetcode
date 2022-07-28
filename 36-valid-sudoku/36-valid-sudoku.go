func isValidSudoku(board [][]byte) bool {
    rows := make([][]int,10)
    cols := make([][]int,10)
    grids := make([][]int,10)
    
    for i := 0 ; i < 9 ; i++ {
        rows[i] = make([]int,10)
        cols[i] = make([]int,10)
        grids[i] = make([]int,10)
    }
    for rowIdx,row := range board {
        for colIdx, key := range row {    
            
            if key == '.'{
                continue
            }
            
            c, _ := strconv.Atoi(string(key))
            gridNum :=(rowIdx / 3 )+ 3 * (colIdx/3)
            
            if rows[rowIdx][c] != 0 || cols[colIdx][c] != 0|| grids[gridNum][c] != 0{
                return false
            }
            
            rows[rowIdx][c] = c
            cols[colIdx][c] = c
            grids[gridNum][c] = c
        }
    }
            
    return true
}