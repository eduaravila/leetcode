func permute(nums []int) [][]int {
    res := [][]int{}
    solution(nums,[]int{},&res)
    return res
}

func solution(nums []int, permutation []int, res *[][]int)bool{
    if len(nums) < 1{
        return true
    }
    
    for i,num := range nums {        
        numsSlice := append([]int{},nums[0:i]...)
        numsSlice = append(numsSlice,nums[i+1:]...)
        if solution(numsSlice, append(permutation,num),res) {
            *res = append(*res,append(permutation,num))
        }
        
    }
    return false
}