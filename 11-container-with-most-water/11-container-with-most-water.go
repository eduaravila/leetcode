func getMin(a,b int) int{
    if a < b { 
        return a
    }
    return b    
}

func getMax(a,b int) int{
    if a > b {
        return a
    }
    return b
}
func maxArea(height []int) int {
    l:= 0 
    r := len(height) -1 
    var res int
    for l < r { 
        width := r-l 
        res = getMax(res, width * getMin(height[l], height[r]))
        if height[l] > height[r]{
            r--            
        }else{
            l++
        }
    }    
    return res
}