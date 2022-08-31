func combinationSum(candidates []int, target int) [][]int {
    res := [][]int{}
    
    solution(candidates, 0, target, []int{}, &res)
    return res
}

func solution(candidates []int,candidatesI, current int, subset []int, res *[][]int) bool{    
    if current == 0 {
         return true
    }
    for i := candidatesI ; i < len(candidates) ; i++ {
        candidate := candidates[i]
        if candidate > current{// continue because input is not sorted
            continue
        }
        subset = append([]int{candidate}, subset...)
        if solution(candidates, i,current-candidate ,subset, res) {
            *res  = append(*res, subset)
        }
        subset = append([]int{},subset[1:]...)
    }
    return false
}