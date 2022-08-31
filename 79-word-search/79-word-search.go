func exist(board [][]byte, word string) bool {       
    
    for x := range board{
        for y := range board[x]{                        
            if solution(board,make(map[string]bool),x,y,0,word){
                return true
            }
        }
    }
    return false
}

func solution(board [][]byte,visited map[string]bool, x,y,current int, word string)bool{
    key := fmt.Sprintf("%d,%d",x,y)
    
    if x < 0 || x >= len(board) || y < 0 || y >= len(board[0]) || visited[key]  || current >= len(word) {
        return false
    }        
    if board[x][y] != byte(word[current]) {
        return false
    }   
    
    if current == len(word)-1 && board[x][y] == byte(word[current]) {
        return true
    }
    visited[key] = true
    if solution(board,visited,x,y+1,current+1,word) || solution(board,visited,x,y-1,current+1,word) || solution(board,visited,x+1,y,current+1,word) || solution(board,visited,x-1,y,current+1,word){
        return true
    }
    visited[key] =false
    return false
}