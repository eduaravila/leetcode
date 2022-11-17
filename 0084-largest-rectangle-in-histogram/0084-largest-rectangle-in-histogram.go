func getMax(a,b int)int{
    if a > b{
        return a
    }
    return b
}

func largestRectangleArea(heights []int) int {
    stack := []int{} // monotonic increasing
    var max , i int
    
    for i <= len(heights){
        var c int
        if i < len(heights){
            c = heights[i]
        }
        
        if len(stack) > 0 && c < heights[stack[len(stack)-1]]{
            element := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            if len(stack) < 1{
                max = getMax(max,heights[element] * i)
            }else{
                max = getMax(max,heights[element] * (i -1 - stack[len(stack)-1]))
            }
            i--
        }else{
            stack = append(stack,i)
        }
        i++
        
    }
    
    return max   
}