func getMin(a,b int)int{
    if a < b{
        return a
    }
    return b
}

const inf = int((^uint(0)) >> 1)

func minPathSum(grid [][]int) int {
    return topDown(grid, len(grid)-1, len(grid[0])-1, make(map[string]int))
}

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

/*
    [[1]] -> 1
    
    
    final point = x = len(grid) y = len(grid[0])
    
    memo[targetcol][targetrow] = grid[targetcol][targetrow] + min(dp(currentrow-1,currentcol),dp(currentrow, currentcol-1)) 
*/