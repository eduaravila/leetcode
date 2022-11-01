func interchangeableRectangles(rectangles [][]int) int64 {
    sums := make(map[float64]int64)
    var res int64
    for _,rectangle := range rectangles{
        width, height := rectangle[0],rectangle[1]
        total := float64(width)/ float64(height)
        if val,e := sums[total];e{
            res += val            
        }
        sums[total]++
        
    }
    
    return res
}