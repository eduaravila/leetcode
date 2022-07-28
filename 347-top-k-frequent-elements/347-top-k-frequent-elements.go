func topKFrequent(nums []int, k int) []int {
    groups := make(map[int]int)
    
    for _,num := range nums {
        groups[num]++
    }
    
    buckets := make([][]int,len(nums)+1)
    for num, qty := range groups {        
        
        buckets[qty] = append(buckets[qty],num)
    }
    res := []int{}
    
    for i:= len(buckets)-1; i >=0 ; i-- {        
        if len(res) >= k {
            return res[:k]
        }
        res = append(res, buckets[i]...)
    } 
    return res[:k]
}