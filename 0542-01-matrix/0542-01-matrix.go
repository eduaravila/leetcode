func updateMatrix(mat [][]int) [][]int {
    
    inf := int(^uint(0)>>1)   
    res := [][]int{}        
    queue := [][]int{}
    for x := range mat{
        res = append(res,make([]int,len(mat[0])))    
        for y := range mat[x]{
            if mat[x][y] == 0{
                res[x][y] = 0
                queue =append(queue,[]int{x,y})
                continue
            }                       
            res[x][y] = inf
        }
    }
    solution(mat,queue, &res)                
    return res
}

func solution(mat, queue [][]int, res *[][]int){            
    dirs := [][]int{{-1,0},{1,0},{0,-1},{0,1}}
    for len(queue) >0{
        element := queue[len(queue)-1]
        queue = queue[:len(queue)-1]  
        for _,dir := range dirs{                                               
            x_n,y_n := element[0] + dir[0], element[1] + dir[1]
            
            if x_n < 0 || y_n < 0 || x_n > len(mat) -1 || y_n > len(mat[0]) - 1 || (*res)[x_n][y_n] <= (*res)[element[0]][element[1]]+1 {
                continue
            }            
            queue = append(queue,[]int{x_n,y_n})
            (*res)[x_n][y_n] = (*res)[element[0]][element[1]]+1
        }
    }
 
}