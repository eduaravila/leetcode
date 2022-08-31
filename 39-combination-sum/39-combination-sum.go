func combinationSum(candidates []int, target int) [][]int {
    res := [][]int{}
    
    solution(candidates,target,[]int{},&res)
    return res
}

func solution(candidates []int, current int,subset []int, res *[][]int) bool{
    if current < 0 {
            return false
    }
    if current == 0 {
         return true
    }        
    for i,candidate := range candidates {
        subset = append([]int{candidate},subset...)
        if solution(append([]int{},candidates[i:]...), (current-candidate),subset,res){
            *res  = append(*res, subset)            
        }
        subset = append([]int{},subset[1:]...)
    }
    
    return false
}