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
    
    row,col := make([]int,m), make([]int,m)
    row[0] = 1
    for y := 0; y < n ; y++{
        for i := 0; i < m ; i++{        
            col[i] = row[i]            
            if i > 0 {
                col[i] += col[i-1]
            }
        }
        row,col = col,row
    }
    return row[m-1]
}