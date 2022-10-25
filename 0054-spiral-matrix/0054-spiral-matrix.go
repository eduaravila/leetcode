func spiralOrder(matrix [][]int) []int {
    directions := [][]int{{0,1},{1,0},{0,-1},{-1,0}}
    lastDirection := 0
    res := []int{}
    
    queue := [][]int{{0,0}}
    for len(queue) >0{
        x,y := queue[len(queue)-1][0],queue[len(queue)-1][1]
        queue = queue[:len(queue)-1]
        
        if matrix[x][y] < -100{
            continue
        }
        
        res =append(res, matrix[x][y])
        matrix[x][y] = -101
        d := 0
        for d < 4{
            directionIndex := (lastDirection + d )% 4            
            direction := directions[directionIndex]
            xn,yn := direction[0],direction[1]
            d++
            if xn+x < 0 || yn+ y < 0 || xn+x > len(matrix)-1 || yn+y > len(matrix[0])-1 || matrix[xn+x][yn+y] < -100{
                continue
            }
            queue = append(queue,[]int{xn+x,yn+y})
            lastDirection = directionIndex
            break
        }
    }
    return res
}