func subarraySum(nums []int, k int) int {
    prefix := make(map[int]int)
    var sum, res int
    prefix[0]++
    for _,num := range nums{
        sum += num
        if val,e := prefix[sum-k];e{
            res+=val
        }
        prefix[sum]++
    }
    
    return res
}