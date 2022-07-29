func longestConsecutive(nums []int) int {
    indexes := make(map[int]bool)
    
    for _,num := range nums {
        indexes[num] = false
    }
    
    maxSequence := 0
    
    for _, num := range nums {
        if _,e :=indexes[num-1]; e {
            continue
        }
        
        current := num
        currentSequence := 0
        _, e := indexes[current];
        for e {            
            currentSequence++
            if currentSequence > maxSequence{
                maxSequence = currentSequence
            }
            current++
            _, e = indexes[current];
        }
    }
    return maxSequence
}