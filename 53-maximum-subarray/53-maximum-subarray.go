func getMax(a,b int)int{
    if a > b {
        return a
    }
    return b
}

func maxSubArray(nums []int) int {
    var total,max,l,r int
    inf := int(^uint(0)>>1)
    max = -inf
    for r < len(nums){
        total += nums[r]
        max = getMax(max,total)
        for total < 0 {
           total -=nums[l]
           l++
        }
        r++
    }
    return max
}