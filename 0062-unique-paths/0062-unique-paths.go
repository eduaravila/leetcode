func getMax(a,b int) int{
    if a > b{
        return a
    }
    return b
}
func uniquePaths(m int, n int) int {
    matrix := [][]int{}
    for i := 0 ; i < m ;i++{
        matrix = append(matrix, make([]int,n))
    }
    
    return solution(matrix,0,0,make(map[string]int))
    
}

func solution(matrix [][]int, x,y int,memo map[string]int) int{
    key := fmt.Sprintf("%d-%d",x,y)
    if val,e := memo[key];e{
        return val
    }
    if x > len(matrix)-1 || y > len(matrix[0])-1  {        
        return 0
    }
    
    if  x == len(matrix)-1 && y == len(matrix[0])-1 {
        return 1
        
    }
    
    total := 0
    matrix[x][y]+=1
    total += solution(matrix,x+1,y,memo)
    total += solution(matrix,x,y+1,memo)        
    memo[key] = total
    return total
    
}