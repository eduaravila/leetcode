func getMax(a, b int)int{
    if a > b{
        return a
    }
    return b
}

func maxProduct(nums []int) int {
    l,r,n := 1,1, len(nums)
    res := nums[0]
    for i := range nums{        
        if l == 0{
            l = 1
        }        
        l *= nums[i]
        if r == 0{
            r = 1
        }
        r *= nums[n-1 - i]        
        res = getMax(res,getMax(l,r))
    }
    
    return res
}