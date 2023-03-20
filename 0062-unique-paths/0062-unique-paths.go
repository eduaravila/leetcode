/*

n, m := len(grid), len(grid[n-1])

if n == 1 and m == 1 
    result = 1


* base case i = n-1 && y == m-1

1,1,1,1,1,1,1
1,2,3,4,5,6,7
1,3,6,10,15,21,28


res[i][y] = matrix[a][b-1] + matrix[a-1][b]

return res[m][n]
*/

func uniquePaths(m int, n int) int {
    if m < 2 || n < 2{
        return 1
    }
    
    res := [][]int{}
    
    for i:= 0; i < m; i++{
        res = append(res, make([]int,n))
    }
    
    res[0][0] = 1
    
    for i := 0; i < m ; i++{
        for y := 0; y < n ; y++{
            if i > 0 {
                res[i][y] += res[i-1][y]
            }
            if y > 0 {
                res[i][y] += res[i][y-1]
            }
            
        }
    }
    
    return res[m-1][n-1]
}