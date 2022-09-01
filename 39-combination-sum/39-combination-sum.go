func combinationSum(candidates []int, target int) [][]int {
    res := [][]int{}
    solution(candidates,0,target,[]int{}, &res)
    return res
}

func solution(candidates []int,current, target int,convination []int, res *[][]int)bool{
    if target == 0 {
        return true
    }
    
    for i := current ;i< len(candidates);i++{
        candidate := candidates[i]
        if candidate > target {
            continue
        }
        
        convination = append(convination,candidate)
        if solution(candidates,i,target-candidate,convination, res){
            *res = append(*res,convination)
        }
        convination = append([]int{},convination[:len(convination)-1]...)
    }
    return false
}