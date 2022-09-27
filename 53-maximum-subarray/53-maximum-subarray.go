func getMax(a,b int)int{
    if a > b {
        return a
    }
    return b
}

func getMin(a,b int)int{
    if a < b {
        return a
    }
    return b
}

func maxSubArray(nums []int) int {
    var max,sum int
        
    max,min,sum := nums[0],nums[0],nums[0]
    for _, num := range nums{
        sum += num
        max = getMax(max, sum- min)
        min = getMin(min, sum)
    }
    return max
}