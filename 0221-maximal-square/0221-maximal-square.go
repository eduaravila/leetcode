func getMax(a,b int)int{
    if a > b {
        return a
    }
    return b
}

func getMin(a,b int)int{
    if a < b {
        return a
    }
    return b
}


func maximalSquare(matrix [][]byte) int {
    var max int
    dp := [][]int{}
    n := len(matrix)
    m := len(matrix[0])
    for i := 0 ; i<= n ; i++ {
        dp = append(dp, make([]int,m+1))
    }
    
    
    
    for x:= 1; x <= n ; x++ {
        for y := 1 ; y <= m; y++{
            top := dp[x-1][y]
            left := dp[x][y-1]
            corner := dp[x-1][y-1]
            if matrix[x-1][y-1] == '0'{
                continue
            }            
            
            dp[x][y] = getMin(top, getMin(left, corner)) + 1
            max = getMax(max, dp[x][y])
        }
    }
    
    
    return max * max
}