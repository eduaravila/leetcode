func recurseUntilLimit(board [][]string, x,y, plusx,plusy int) bool{
    if  x < 0 || x  >= len(board) || y < 0 || y  >= len(board)  {
        return true
    }
    
    if board[x][y] == "Q"{
        return false
    }
    return recurseUntilLimit(board, x+plusx, y + plusy,plusx,plusy)
    
}

func isValid(board [][]string, x,y int)bool{      
    return recurseUntilLimit(board,x,y,0,1) && recurseUntilLimit(board,x,y,0,-1) && recurseUntilLimit(board,x,y,1,0) && recurseUntilLimit(board,x,y,-1,0) && recurseUntilLimit(board,x,y,1,1) && recurseUntilLimit(board,x,y,1,-1) && recurseUntilLimit(board,x,y,-1,-1) && recurseUntilLimit(board,x,y,-1,1)
}


func solveNQueens(n int) [][]string {
    if n < 1 {
        return [][]string{}
    }
    res := [][]string{}
    board := make([][]string,n)
    
    for x := 0 ; x < n ; x++{
        
        for y := 0 ; y < n ; y++{
            board[x] = append(board[x],".")
        }
    }
    
    solution(board,0,n,&res)
    return res
}

func solution(board [][]string,currentx, queens int, res *[][]string) {
    
    if queens <1 {
        boardString := []string{}
        for _,row := range board {
            boardString = append(boardString, strings.Join(row,""))
        }
        
        *res = append(*res, boardString)
        return
    }
    
    for x := currentx; x < len(board) ; x++{
        for y := 0; y < len(board) ; y++{
            
            if isValid(board,x,y){                    
                board[x][y] = "Q" 
                
                solution(board,x, queens-1,res)
                board[x][y] = "."
            }
            
        }
    }
}