func orangesRotting(grid [][]int) int {
    // get rotting oranges
    queue := [][]int{}
    
    for x := range grid{
        for y, space := range grid[x]{
            if space == 2 {
                queue= append(queue, []int{x,y})
            }
        }
    }
    
    var minutes int
    for len(queue) > 0 {
        newQ := [][]int{}

        for _,space := range queue {

            left := []int{space[0] -1, space[1]}
            right := []int{space[0] +1, space[1]}
            top := []int{space[0], space[1]-1}
            bottom := []int{space[0], space[1]+1}
            
            if isInLimits(grid,left) && grid[left[0]][left[1]] == 1 {
                newQ = append(newQ, left)
                grid[left[0]][left[1]]= 2 
            }
            
            if isInLimits(grid,right) && grid[right[0]][right[1]] == 1{
                newQ = append(newQ, right)
                grid[right[0]][right[1]] = 2
            }
            
            if isInLimits(grid,top) && grid[top[0]][top[1]] == 1{
                newQ = append(newQ, top)
                grid[top[0]][top[1]]= 2
            }
            
            if isInLimits(grid,bottom) && grid[bottom[0]][bottom[1]] == 1{
                newQ = append(newQ, bottom)
                grid[bottom[0]][bottom[1]] =2 
            }
        }
        if len(newQ) > 0{
            minutes++
        }
        queue = newQ
    }
    
    
    for x := range grid{
        for _, space := range grid[x]{
            if space == 1 {
               return -1
            }
        }
    }
    
    return minutes    
    // iterare through the roten oranges and get neigborh 
}

/*
2 1 1 
0 1 1
1 0 1
*/
func isInLimits(grid [][]int, coordinate []int)bool{
    if coordinate[0] < 0 || coordinate[1] < 0 || coordinate[0] > len(grid)-1 || coordinate[1] > len(grid[0])-1 {
        return false 
    }
    return true
}