func getMax(a,b int)int{
    if a > b{
        return a
    }
    return b
}

func largestRectangleArea(heights []int) int {
    var max, i int
    stack := []int{}
    
    for i <= len(heights) {                
        var c int
        if i < len(heights){
            c = heights[i]
        }
        n := len(stack)-1
        if len(stack) > 0 && c < heights[stack[n]] {
            e := stack[n]
            stack = stack[:n]
            if len(stack) < 1{
                max = getMax(max, heights[e] * (i))
            }else{
                max= getMax(max, heights[e] * ((i - 1) - (stack[len(stack)-1])))
            }            
            i--
         }else{        
            stack = append(stack,i)
         }
        i++
    }      
    
    return max
}          


