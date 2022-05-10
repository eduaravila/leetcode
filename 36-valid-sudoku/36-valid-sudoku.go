import 
(
    "strconv"
)

func isDuplicated(arr []byte, val byte) bool{    
    for _, current := range arr {
        if current == val {
            return true
        }
    }
    return false
}

func getSubboardKey(col, row int) string{
    colK := strconv.Itoa(int(col)/3)
    rowK := strconv.Itoa(int(row)/3)
    return colK + ","+ rowK 
}

func isValidSudoku(board [][]byte) bool {
    rows := make(map[int][]byte)
    cols := make(map[int][]byte)
    subs := make(map[string][]byte)
    
    for idxRow,row :=  range board{
        for idxCol, val := range row{ 
            if val == '.'{
                continue
            }
            if isDuplicated(rows[idxRow], val) || isDuplicated(cols[idxCol], val) || isDuplicated(subs[getSubboardKey(idxCol, idxRow)], val){
                return false
            }
            rows[idxRow] = append(rows[idxRow],val)
            cols[idxCol] = append(cols[idxCol],val)
            subs[getSubboardKey(idxCol,idxRow)] = append(subs[getSubboardKey(idxCol,idxRow)], val)
        }
    }
    return true
}