func getMax(a,b float64) float64{
    if a > b {
        return a
    }
    return b
}

func findMaxAverage(nums []int, k int) float64 {
    l, r := 0,0
    inf := int(^uint(0)>>1)
    max := float64(-inf)
    var total float64
    for r < len(nums){  
        for r < len(nums) && r - l +1 <= k{
            total += float64(nums[r])
            r++
        }
        max = getMax(max, total / float64(k))
        total -= float64(nums[l])
        l++        
    }
    
    return max
}