func getMax(a,b int)int{
    if a > b {
        return a
    }
    return b
}

func mergeTriplets(triplets [][]int, target []int) bool {
    
    max := make([]int,3)
    
    for _,triplet := range triplets{
        
        current := 0
        for j:=0 ; j < 3 ; j++{            
            value := triplet[current]
            if value > target[j]{
                break
            }
            current++
        }
        if current == 3{
            for i, val := range triplet{
                max[i] = getMax(val,max[i])
            }
        }
        
    }
    
    for i,value := range target{
        if value != max[i]{
            return false
        }
    }
    return true
}