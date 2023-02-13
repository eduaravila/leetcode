func getMax(a,b int) int{
    if a > b{
        return a
    }
    return b
}

func maxFrequency(nums []int, k int) int {    
    var max,l,r,sum int
    sort.Slice(nums,func(a,b int) bool { return nums[a] < nums[b]})
    
    for r < len(nums) {
        sum += nums[r]
        window := ((r-l) +1) * nums[r]
        if window > (sum+k){
            sum -= nums[l]
            l++
        }
        max = getMax(max, r-l+1)
        r++
    }
    return max
}