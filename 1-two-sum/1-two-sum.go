func twoSum(nums []int, target int) []int {
    elements := make(map[int]int,0)
    for idx,val := range nums {
        rest := target - val
        if val,e := elements[rest]; e {
            return []int{idx,val}
        }
        elements[val] = idx
    }   
    return []int{}
}