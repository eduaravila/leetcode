func subsets(nums []int) [][]int {
    res := [][]int{}
    solution(nums,[]int{},&res)
    return append(res,[]int{})
}

func solution(nums, current []int,res *[][]int) {
    if len(nums) < 1{
        return 
    }
    
   
    
    for i,num := range nums {                   
        current = append(current,num)
        *res = append(*res, current)
        newNums := append([]int{},nums[i+1:]...)
        solution(newNums,current,res)                
        current = append([]int{},current[:len(current)-1]...)
    }    
    
}