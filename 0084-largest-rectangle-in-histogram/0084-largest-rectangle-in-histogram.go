func getMax(a,b int)int{
    if a > b{
        return a
    }
    return b
}
func largestRectangleArea(heights []int) int {
    

    stack :=[]int{}
    var max int
    var i int
    for i <= len(heights){   
        var h int
        if i < len(heights){
            h = heights[i]
        }
        
        if  len(stack) > 0 && h < heights[stack[len(stack)-1]]{
            tp := heights[stack[len(stack)-1]]
            stack = stack[:len(stack)-1]
            if len(stack) <1{
                max = getMax(max, tp * (i))                
            }else{
                max = getMax(max, tp * (i - 1 - stack[len(stack)-1] ))                
            }
            
            i--
            
        }else{
            stack = append(stack,i)    
        }
        i++
        
    }    
    
    return max
} 