func updateMatrix(mat [][]int) [][]int {
    
    inf := int(^uint(0)>>1)   
    
    queue := [][]int{}
    for x := range mat{
        
        for y := range mat[x]{
            if mat[x][y] != 0{ 
                mat[x][y] = inf
                continue
            }   
            queue =append(queue,[]int{x,y})
            
        }
    }
    return solution(mat,queue)
    
}

func solution(mat, queue [][]int)[][]int{            
    dirs := [][]int{{-1,0},{1,0},{0,-1},{0,1}}
    for len(queue) >0{
        l := len(queue)
        for _,element := range queue{        
            for _,dir := range dirs{                                               
                x_n,y_n := element[0] + dir[0], element[1] + dir[1]
                if x_n < 0 || y_n < 0 || x_n > len(mat) -1 || y_n > len(mat[0]) - 1 || mat[x_n][y_n] <= mat[element[0]][element[1]]+1 {
                    continue
                }            
                queue = append(queue,[]int{x_n,y_n})
                mat[x_n][y_n] = mat[element[0]][element[1]]+1
            }
        }
        queue = queue[l:]
    }
    return mat
}