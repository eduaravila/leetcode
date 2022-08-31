func subsetsWithDup(nums []int) [][]int {    
    res := [][]int{}
    sort.Ints(nums)
    solution(nums,[]int{},&res)
    return append(res,[]int{})
}

func solution(nums ,subset []int, res *[][]int)bool{
    if len(nums) < 1{
        return true
    }
     
    for i,num := range nums {
        if i > 0 && nums[i] == nums[i-1]{
            continue
        }
        subset = append(subset,num)      
        *res = append(*res, subset)
        newS := append([]int{},nums[i+1:]...)
        solution(newS,subset,res)
        subset = append([]int{},subset[:len(subset)-1]...)
    }
    
    return false
}