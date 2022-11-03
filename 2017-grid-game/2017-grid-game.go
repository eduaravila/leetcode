func getMax(a,b int)int{
    if a > b{
        return a
    }
    return b
}

func getMin(a,b int)int{
    if a < b{
        return a
    }
    return b
}

func gridGame(grid [][]int) int64 {
    top,bottom := make([]int,len(grid[0])), make([]int,len(grid[0]))
    n := len(grid[0])-1
    var t,b int
    for i := range grid[0]{        
        top[n-i] = t
        bottom[i] = b
        t += grid[0][n-i]
        b += grid[1][i]
    }
    var max int
    res := int(^uint(0)>>1)
    for i := range grid[0]{        
        max = getMax(top[i],bottom[i])
        res =getMin(max, res)
    }    
    return int64(res)    
}