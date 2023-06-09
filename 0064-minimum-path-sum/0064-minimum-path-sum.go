func getMin(a,b int)int{
    if a < b{
        return a
    }
    return b
}

const inf = int((^uint(0)) >> 1)

func minPathSum(grid [][]int) int {
    table := [][]int{}
    m, n := len(grid[0]), len(grid)
    
    for i := 0 ; i <= n ;i++ {     
        infRow := []int{}
        for x := 0 ; x <= m ;x++ {     
            infRow = append(infRow,inf)
        }
        table = append(table,infRow  )
    }
    
    table[n-1][m-1] = 0
    table[n][m-1] = 0
    table[n-1][m] = 0
    for x:= n-1 ; x >-1 ; x--{        
        for y := m-1; y>-1;y--{
            
            right := table[x][y+1]
            bottom := table[x+1][y]
            
            table[x][y] = grid[x][y] + getMin(right,bottom)
        }
    }
    
    return table[0][0]
}


/*

func topDown(grid[][]int, col, row int, memo map[string]int ) int{
    key := fmt.Sprintf("%d-%d",col,row)
    
    if val, e := memo[key];e{
        return val
    }
    if col == 0 && row == 0{
        return grid[col][row]
    }
    if col < 0 || row < 0{
        return inf
    }   
    
    min := getMin(topDown(grid,col-1, row, memo), topDown(grid,col,row-1, memo))
    memo[key] =min + grid[col][row] 
    return memo[key]
}

    [[1]] -> 1
    
    
    final point = x = len(grid) y = len(grid[0])
    
    memo[targetcol][targetrow] = grid[targetcol][targetrow] + min(dp(currentrow-1,currentcol),dp(currentrow, currentcol-1)) 
    
    
     [[1,3,1],[1,5,1],[4,2,1]]
     
     [7,6,3]
     [5,7,2]
     [4,2,1]
*/