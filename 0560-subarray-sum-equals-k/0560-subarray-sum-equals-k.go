func subarraySum(nums []int, k int) int {
    prefixes := make(map[int]int)
    var res,sum int
    prefixes[0]++
    for _, num := range nums{
        sum += num
        if val, e := prefixes[sum-k];e{
            res+=val
        }
        prefixes[sum]++
    }
    
    return res
}