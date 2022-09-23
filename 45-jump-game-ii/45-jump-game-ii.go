func getMax(a,b int)int{
    if a > b { 
        return a
    }
    return b
}

func jump(nums []int) int {
    if len(nums) < 2{
        return 0
    }
    
    var furtst, end , i ,res int
    
    for i < len(nums) - 1 {
        furtst = getMax(furtst, nums[i] + i)
        if i == end{
            res++
            end = furtst
        }
        i++
    }
    return res
}