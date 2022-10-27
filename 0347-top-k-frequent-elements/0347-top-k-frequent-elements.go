func topKFrequent(nums []int, k int) []int {
    counter := make(map[int]int)
    for _, num := range nums {
        counter[num]++
    }
    
    buckets := make([][]int,len(nums)+1)
    
    for num,times := range counter{
        buckets[times] = append(buckets[times], num)
    }
    res := []int{}
    i := len(nums)
    for i >= 0 && k > 0{
        for _,num := range buckets[i]{
            res =append(res,num)
            k--    
        }
        i--
    }
    return res
}