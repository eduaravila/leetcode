func getMax(a,b int)int{
    if a > b{
        return a
    }
    return b
}

func rob(nums []int) int {
    var max int
    sums:= append([]int{},nums...)
    for i := range nums{
        if i > 1{
            sums[i] = getMax((nums[i]+(sums[i-1] - nums[i-1])), nums[i]+sums[i-2])
        }
        max = getMax(max,sums[i])
    }    
    return max
}

// [2,7,9,3,1]
// if i > 1 {
// nums[i] += nums[i-2]
//}
// [2,7,11,10,0]
// 10
// update max accordingly