func updateMatrix(mat [][]int) [][]int {
    distances := [][]int{}
    inf := int(^uint(0)>>1)
    queue := [][]int{}
    for y := range mat {
        distances = append(distances, make([]int,len(mat[0])))
        for x := range mat[y]{
            if mat[y][x] != 0 {
                distances[y][x] = inf
            }else{
                queue = append(queue, []int{y,x})
            }
        }
    }
    
    
    coordinates := [][]int{{1,0},{-1,0},{0,1},{0,-1}}
    for len(queue) > 0 {
        x,y := queue[0][0],queue[0][1]
        queue = queue[1:]
        if mat[x][y] < 0 {
            continue
        }
        for _, coordinate := range coordinates {
            nx,ny := x + coordinate[0], y + coordinate[1]
            // outside boundaries or the jump number is bigger than the current
            if  ny < 0 || nx < 0 || nx > len(mat) -1 || ny > len(mat[0])-1 || distances[x][y] + 1 >= distances[nx][ny]{
                continue
            }
            queue = append(queue,[]int{nx,ny})
            mat[x][y] = -1
            distances[nx][ny] = distances[x][y] + 1
        }
    }
    return distances    
}