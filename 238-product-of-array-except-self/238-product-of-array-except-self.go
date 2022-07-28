func productExceptSelf(nums []int) []int {
    res := make([]int,0)
    for range nums{
        res = append(res,1)
    }
    prex := 1 
    for i:= range res{
        res[i] = prex
        prex *= nums[i]
    }
    
    posx := 1
    for i := len(nums)-1 ; i >= 0 ; i--{
        res[i] *= posx
        posx *= nums[i]
    }
    return res 
}