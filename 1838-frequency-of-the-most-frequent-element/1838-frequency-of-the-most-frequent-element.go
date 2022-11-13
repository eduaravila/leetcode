func getMax(a,b int)int{
    if a > b {
        return a
    }
    return b
}

func maxFrequency(nums []int, k int) int {
    sort.Ints(nums)
    var res,sum,l,r int
    
    for r < len(nums){        
        sum += nums[r]        
        window := (r-l+1) * nums[r]        
        if window > (sum + k){            
            sum -= nums[l]
            l++
        }
        res = getMax(res,r-l+1)
        r++
    }
    return res
}