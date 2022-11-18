func sortedSquares(nums []int) []int {
    res := make([]int,len(nums))
    
    l,r := 0,len(nums)-1
    i := r
    for i >= 0{
        if math.Abs(float64(nums[l])) > math.Abs(float64(nums[r])){
            res[i] = nums[l] *  nums[l]
            l++
        }else{
            res[i] =  nums[r] *  nums[r]
            r--
        }
        i--        
    }
    return res
}