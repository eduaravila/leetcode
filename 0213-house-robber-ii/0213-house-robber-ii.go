func getMax(a,b int)int{
    if a > b{
        return a
    }
    return b
}

func rob(nums []int) int {        
    if len(nums) < 2{
        return nums[0]
    } 
    return getMax(robStuff(nums[:len(nums)-1]),robStuff(nums[1:]) )
}


func robStuff(nums []int)int{
    var max int
    sums := make([]int,len(nums))
    for i , num := range nums{        
        if i > 1 {
            sums[i] = getMax(num + (sums[i-1] - nums[i-1]),num + sums[i-2])
        }else{
            sums[i] = num
        }
        max = getMax(sums[i],max)
    }
    
    return max
}