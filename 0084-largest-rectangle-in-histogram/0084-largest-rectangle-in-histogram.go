func getMax(a,b int)int{
    if a> b{
        return a
    }
    return b
}

func largestRectangleArea(heights []int) int {
    var max int
    stack:= []int{}
    
    var r int
    for r <= len(heights){
        val := 0 
        if r < len(heights){
            val = heights[r] 
        }     
        if len(stack) > 0 && val < heights[stack[len(stack)-1]]{                
                colIndex := stack[len(stack)-1]
                colVal := heights[colIndex]      
                stack = stack[:len(stack)-1]
                if len(stack) < 1{
                    max = getMax(colVal * r, max)
                }else{
                    w := (r-1) - stack[len(stack)-1]
                    max = getMax(max, colVal * w)
                }
            r--
        }else{
            stack =append(stack, r)
        }
        r++
    }
    return max
}