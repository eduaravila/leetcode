func spiralOrder(matrix [][]int) []int {
    directions := [][]int{{0,1},{1,0},{0,-1},{-1,0}}
    var direction int
    visited := [][]bool{}
    for  range  matrix{
        visited = append(visited, make([]bool,len(matrix[0])))
    }
    
    stack := [][]int{{0,0}}
    res := []int{}
    for len(stack) > 0 {        
        x,y := stack[len(stack)-1][0],stack[len(stack)-1][1]
        stack = stack[:len(stack)-1]
        
        visited[x][y] = true
        res = append(res, matrix[x][y])
        var c_dir int
        x_new, y_new := x+directions[direction][0], y+directions[direction][1]
        for c_dir < len(directions) && (x_new < 0 || y_new < 0 || x_new > len(matrix)-1 || y_new > len(matrix[0])-1 || visited[x_new][y_new] )   {
            direction = (direction + 1) % len(directions)
            x_new, y_new = x+directions[direction][0], y+directions[direction][1]
            c_dir++
        }
        
        if c_dir > len(directions)-1{
            break
        }
        stack = append(stack, []int{x_new,y_new})
    }
    return res
}