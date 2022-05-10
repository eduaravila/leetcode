
const 
(
    COL_SIZE = 3
    ROW_SIZE = 3
    BOARD_TOTAL = 9
)

func convertToInt(val byte) int{
    if val >= '1' && val <= '9' {
        return int(val)
    }
    return 0
}


func isValidSudoku(board [][]byte) bool {
    
    unique := make(map[int]bool)
    
    for _ ,row := range board {
        for _, val := range row {
            value := convertToInt(val)
            if value == 0 {
                continue
            }
            if _,e := unique[value];e {
                return false
            }
            unique[value] = true
        }
        unique = make(map[int]bool)        
    }
    
        
    for idxcol, col := range board {
        for idxrow := range col {
            value := convertToInt(board[idxrow][idxcol])
            
            if value == 0 {
                continue
            }
            if _,e := unique[value];e {
                return false
            }
            unique[value] = true
        }
        unique = make(map[int]bool)        
    }
    
    
    for n_board := 0 ; n_board < BOARD_TOTAL; n_board++{
        for field := 0 ; field < COL_SIZE * ROW_SIZE; field++{
            col := (field % COL_SIZE) + ((n_board % COL_SIZE) * COL_SIZE) 
            row := field / ROW_SIZE + ((n_board / ROW_SIZE) * ROW_SIZE)
            value := convertToInt(board[row][col])
            if value == 0 {
                continue
            }
            if _,e := unique[value];e {
                return false
            }
            unique[value] = true
        }     
        unique = make(map[int]bool)        

    }
    
    return true
}