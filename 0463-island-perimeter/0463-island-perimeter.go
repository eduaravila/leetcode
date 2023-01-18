func islandPerimeter(grid [][]int) int {
    var res int
    
    for x := range grid{
        for y, val := range grid[x]{
            if val == 1{
                res += 4 -numberOfNeighbors(x,y,grid)
            }
        }
    }
    
    return res
}

func numberOfNeighbors(x,y int, grid[][]int) int{
    top,bottom,left,right := x-1,x+1,y-1,y+1
    var res int
    if top >= 0 && grid[top][y] == 1 {
        res++
    }
    if  bottom < len(grid) && grid[bottom][y] == 1  {
        res++
    }
    
    if left >= 0 && grid[x][left] == 1{
        res ++
    }
    
    if right < len(grid[len(grid)-1]) && grid[x][right] ==1{
        res++
    }
    return res
}