func topKFrequent(nums []int, k int) []int {
    counter := make(map[int]int)
    
    for _,num := range nums {
        counter[num]++
    }
    groups := make([][]int,len(nums)+1)
    
    for num,freq := range counter{
        groups[freq] = append(groups[freq], num)
    }
    
    res := make([]int,len(nums))
    
    for _,group := range groups{
        res = append(res, group...)
    }
    
    return res[len(res)-k:]
}