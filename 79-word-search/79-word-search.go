func exist(board [][]byte, word string) bool {       
    
    for x := range board{
        for y := range board[x]{
            
            if solution(board,x,y,0,word){
                return true
            }
        }
    }
    return false
}

func solution(board [][]byte, x,y,current int, word string)bool{
    
    if x < 0 || x >= len(board) || y < 0 || y >= len(board[0])  || current >= len(word) || board[x][y] == '*'{
        return false
    }        
    if board[x][y] != byte(word[current]) {
        return false
    }   
    
    if current == len(word)-1 && board[x][y] == byte(word[current]) {
        return true
    }
    
    board[x][y] = '*'
    if solution(board,x,y+1,current+1,word) || solution(board,x,y-1,current+1,word) || solution(board,x+1,y,current+1,word) || solution(board,x-1,y,current+1,word){
        return true
    }
    board[x][y] = word[current]
    return false
}