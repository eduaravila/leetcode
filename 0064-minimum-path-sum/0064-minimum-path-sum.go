/*

base case 
    x and y == m, n


to get grid[m][n] we need grid[m][n] steps


if the is no value available to the right choose 

[i,3,1]
[i,i,0]


[i,i,i]
[i,i,i]
[i,i,i]


*/


func getInfArrOfSize(size int)[]int{
    inf := int(^uint(0)>>1)
    res := make([]int,size)
    
    for i := range res{
        res[i] = inf
    }
    return res
}

func getMin(a,b int)int{
    if a < b{
        return a
    }
    return b
}

func minPathSum(grid [][]int) int {
    m,n := len(grid),len(grid[0])
    res := getInfArrOfSize(n)
    res[n-1] = 0
    
    for i := m-1; i > -1 ; i--{
        dp := getInfArrOfSize(n)
        for x := n-1; x > -1; x--{
            if x < n-1{
                dp[x] = getMin(dp[x+1], dp[x])
            }
            dp[x] = getMin(dp[x], res[x]) + grid[i][x]
        }
        res = dp
    }
    
    return res[0]
}