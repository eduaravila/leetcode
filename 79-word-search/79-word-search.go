func exist(board [][]byte, word string) bool {       
    visited := [][]bool{}
    for range board{
        visited = append(visited,make([]bool,len(board[0])))
    }
    for x := range board{
        for y := range board[x]{
            if visited[x][y]{
                continue
            }
            if solution(board,visited,x,y,0,word){
                return true
            }
        }
    }
    return false
}

func solution(board [][]byte,visited [][]bool, x,y,current int, word string)bool{
    
    if x < 0 || x >= len(board) || y < 0 || y >= len(board[0]) || visited[x][y]  || current >= len(word) {
        return false
    }        
    if board[x][y] != byte(word[current]) {
        return false
    }   
    
    if current == len(word)-1 && board[x][y] == byte(word[current]) {
        return true
    }
    visited[x][y] = true
    if solution(board,visited,x,y+1,current+1,word) || solution(board,visited,x,y-1,current+1,word) || solution(board,visited,x+1,y,current+1,word) || solution(board,visited,x-1,y,current+1,word){
        return true
    }
    visited[x][y] =false
    return false
}