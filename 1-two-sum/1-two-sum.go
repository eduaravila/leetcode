func twoSum(nums []int, target int) []int {
    elements := make(map[int]int)
    for idx,val := range nums {
        rest := target - val
        if _,e := elements[rest]; e {
            return []int{idx,elements[rest]}
        }
        elements[val] = idx
    }   
    return []int{}
}