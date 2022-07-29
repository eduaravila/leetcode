func longestConsecutive(nums []int) int {
    indexes := make(map[int]bool)
    
    for _,num := range nums {
        indexes[num] = false
    }
    
    maxSequence := 0
    
    for num, visited := range indexes {
        if visited {
            continue
        }
        current := num
        currentSequence := 0
        _, e := indexes[current];
        for e {
            indexes[current]= true
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