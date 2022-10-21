func getMax(a,b int)int{
    if a > b {
        return a
    }
    return b
}

func getMin(a,b int) int{
    if a < b {
        return a
    }
    return b
}

func maxArea(height []int) int {
    var max int
    l,r:= 0 ,len(height)-1
    for l < r{
        min := getMin(height[l],height[r])
        max = getMax(min * (r - l), max)
        if height[l] < height[r]{
            l++
        }else if height[r] <= height[l]{
            r--
        }
    }
    return max
}