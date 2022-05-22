func twoSum(nums []int, target int) []int {
    counter := make(map[int]int)
    
    for idx,num := range nums {
        counter[num] =idx // we just need to save the latest index
    }
    
    for idxc, num := range nums{
        remaining := target - num 
        
        if idx, e := counter[remaining]; e && idxc != idx {
            return []int{idxc, counter[remaining]}
        }
    }
    return []int{}
}