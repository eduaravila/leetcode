func subsets(nums []int) [][]int {
    res := [][]int{}
    visited := make(map[string]bool)
    solution(nums,visited,&res)
    return append(res,[]int{})
}

func solution(nums []int,visited map[string]bool,res *[][]int) {
    if len(nums) < 1{
        return 
    }
    
    if visited[fmt.Sprintf("%v",nums)]{
        return
    }
    *res = append(*res, nums)
    for i := range nums {           
        newNums := append([]int{},nums[0:i]...)
        newNums = append(newNums,nums[i+1:]...)      
        solution(newNums,visited,res)                
    }
    visited[fmt.Sprintf("%v",nums)] = true
    
}