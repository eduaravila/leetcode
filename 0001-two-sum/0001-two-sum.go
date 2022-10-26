func twoSum(nums []int, target int) []int {
    mapping := make(map[int]int)
    
    for i, num := range nums{
        if x,e := mapping[target-num];e{
            return []int{i,x}
        }
        mapping[num] =i
    }
    return []int{}
}