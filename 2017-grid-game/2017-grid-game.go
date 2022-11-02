func getMin(a,b int)int{
    if a <  b {
        return a
    }
    return b
}

func getMax(a,b int)int{
    if a > b {
        return a
    }
    return b
}

func gridGame(grid [][]int) int64 {
    preffixes :=[][]int{}    
    
    for i := range grid{
        preffixes = append(preffixes,[]int{})
        var sum int
        for y := 0 ; y < len(grid[0]) ; y++{
            sum += (grid[i][y])
            preffixes[i] = append(preffixes[i] , sum)
        }
    }
    
    var max int
    res := int(^uint(0)>>1)
    for i := range grid[0] {
        top,bottom := preffixes[0][len(preffixes[0])-1] - preffixes[0][i] , 0
        if i > 0 {
            bottom = preffixes[1][i-1]
        }
        max = getMax(top, bottom)
        res = getMin(res,max)
    }
    
    return int64(res)
}

// 2 7 11
// 1 6 7

