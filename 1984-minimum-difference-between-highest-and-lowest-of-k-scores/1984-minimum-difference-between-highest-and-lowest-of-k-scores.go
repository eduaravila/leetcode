
func getMin(a,b int)int{
    if a < b {
        return a
    }
    return b
}


func minimumDifference(nums []int, k int) int {
    if k < 2 {
        return 0
    }
    sort.Ints(nums)
    min := int(^uint(0)>>1)
    for l, r := len(nums)-k, len(nums)-1; l >=0 ; l,r = l-1,r-1{
            min = getMin(min, nums[r] - nums[l])
        }
    
    return min
}