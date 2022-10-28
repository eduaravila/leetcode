func getMax(a,b int)int{
    if a > b{
        return a
    }
    return b
}

func largestOverlap(img1 [][]int, img2 [][]int) int {
        
    positives1 := [][]int{}
    positives2 := [][]int{}        
    for x := range img1{
        for y := range img1[x]{
            if img1[x][y] == 1{
                positives1 = append(positives1,[]int{x,y})
            }
            if img2[x][y] == 1{
                positives2 = append(positives2 ,[]int{x,y})
            }            
        }
    }
    
    counter := make(map[string]int)
    for _, a := range positives1{
        for _, b := range positives2{
            key := fmt.Sprintf("%d-%d",a[0]-b[0], a[1]-b[1])
            counter[key]++
        }
    }
    var max int
    for _,val := range counter{
        max = getMax(max,val)
    }
    return max
}