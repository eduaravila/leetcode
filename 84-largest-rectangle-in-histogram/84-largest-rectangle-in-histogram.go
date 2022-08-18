func getMax(a,b int) int{
    if a > b {
        return a
    }
    return b
}

func largestRectangleArea(heights []int) int {
    stack := [][]int{}
    var maxV int
    for i,val := range heights {
        startI := i
        for len(stack) > 0 && stack[len(stack)-1][1] > val {
            maxV = getMax(maxV, stack[len(stack)-1][1] * (i - stack[len(stack)-1][0]))
            startI = stack[len(stack)-1][0]            
            stack = stack[:len(stack)-1]            
        }
        stack = append(stack,[]int{startI,val})
    }
    
    for _, current := range stack { // this should sorted in ascending order
        maxV = getMax(maxV, current[1] * (len(heights)- current[0]))
    }
    return maxV
} 