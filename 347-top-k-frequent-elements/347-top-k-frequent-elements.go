func topKFrequent(nums []int, k int) []int {
    groups := make(map[int]int)
    
    for _,num := range nums {
        groups[num]++
    }
    
    buckets := make([][]int,len(nums))
    for num, qty := range groups {        
        position := (len(nums) -1)/ qty
        buckets[position] = append(buckets[position],num)
    }
    res := []int{}
    
    for i:= 0 ; i < len(buckets) ; i++ {
        
        if len(res) >= k {
            fmt.Println(res)
            return res[:k]
        }
        res = append(res, buckets[i]...)
    } 
    return res[:k]
}