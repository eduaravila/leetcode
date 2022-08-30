func permuteUnique(nums []int) [][]int {
    res := [][]int{}
    sort.Ints(nums)
    solution(nums,[]int{},&res)
    return res
}

func solution(nums, current[]int, res *[][]int)bool{
    if len(nums) < 1{
        return true
    }
    
    
    for i,num := range nums{
        if i > 0 && nums[i] == nums[i-1]{
            continue
        }
        newNums := append([]int{},nums[i+1:]...)
        newNums = append(newNums,nums[0:i]...)        
        if solution(newNums,append([]int{num},current...),res){
            *res = append(*res,append(current,num))
        }               
    }
    return false
}