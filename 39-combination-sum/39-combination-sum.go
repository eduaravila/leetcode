func combinationSum(candidates []int, target int) [][]int {
    res := [][]int{}
    solution(candidates, 0, target, []int{}, &res)
    return res
}

func solution(candidates []int,candidatesI, current int, subset []int, res *[][]int) bool{
    if current < 0 || candidatesI >= len(candidates){
        return false
    }
    if current == 0 {
         return true
    }    
    candidate := candidates[candidatesI]
    subset = append([]int{candidate}, subset...)
    if solution(candidates, candidatesI,current-candidate ,subset, res) {
        *res  = append(*res, subset)
    }    
    subset = append([]int{},subset[1:]...)
    solution(candidates, candidatesI+1,current ,subset, res) 
    return false
}