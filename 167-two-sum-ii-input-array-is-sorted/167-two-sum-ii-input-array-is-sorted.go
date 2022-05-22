func twoSum(numbers []int, target int) []int {
    l := 0
    r := 1
    res := []int{}
    for r < len(numbers) {
        remaining := target - numbers[l]
        
        for r < len(numbers) && numbers[r] <= remaining{
            if numbers[r] == remaining{                            
                return []int{l+1, r+1}
            }
            r++
        }        
        res = []int{}
        l++
        r = l+1
    }
    return res
}