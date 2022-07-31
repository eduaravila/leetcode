
func maxArea(height []int) int {    
    l,r := 0,len(height) -1
    var currentAmount, maxAmount int
    
    for l < r {                
        if height[l] < height[r]{
            currentAmount = height[l] * (r - l)
            l++
        }else{
            currentAmount = height[r] * (r - l)
            r--
        }
        if currentAmount > maxAmount { 
            maxAmount = currentAmount
        }
    }
    return maxAmount
}