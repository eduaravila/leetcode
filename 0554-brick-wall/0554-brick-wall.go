func getMax(a,b int)int{
    if a > b{
        return a
    }
    return b
}

func leastBricks(wall [][]int) int {
    
    edges_per_row := make(map[int]int)
    var total_rows int
    for _,col := range wall[0]{
        total_rows+=col
    }
    for _, row := range wall{
        var edge int
        for _, col := range row{                        
            edge += col                                    
            edges_per_row[edge]++
        }
    }
    var max int
    
    
    for edge,edges := range edges_per_row{
        if edge == total_rows{
            continue
        }
        max = getMax(max,edges)
    }
    
    if max < 1{
        return len(wall)
    }
    return len(wall)- max
}