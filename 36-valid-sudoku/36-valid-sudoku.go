func isValidSudoku(board [][]byte) bool {
    rows := make(map[string]bool)
    cols := make(map[string]bool)
    
    for rowIdx,row := range board {
        for colIdx, key := range row {    
            keyCol := fmt.Sprintf("%d-%s",colIdx,string(key))
            keyRow := fmt.Sprintf("%d-%s",rowIdx,string(key))
            
            if e := cols[keyCol]; e{                
                return false
            }
            if e := rows[keyRow]; e{                
                return false
            }
            
            if key != '.'{                
                cols[keyCol] = true
                rows[keyRow] = true
            }
            
            
        }
    }
    
    grids := make(map[string]bool)    
     for rowIdx,row := range board {
        for colIdx,val := range row {                      
            key := fmt.Sprintf("%d-%d-%s",(rowIdx / 3),(colIdx / 3),string(val))            
            
            if e := grids[key]; e{                
                return false
            }
            if val != '.'{                
                grids[key] = true
            }
        }
    }
    
    return true
}