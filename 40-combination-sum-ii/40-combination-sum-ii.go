func combinationSum2(candidates []int, target int) [][]int {
    res := [][]int{}
    sort.Ints(candidates)
    solution(candidates, target,[]int{}, &res)
    return res
}

func solution(candidates[]int,target int, convination []int, res *[][]int)bool{
    if target < 0 {
        return false
    }
    if target == 0 {
        return true
    }
    
    
    for i,candidate := range candidates{
        if candidate > target {
            continue
        }
        if i > 0 && candidates[i] == candidates[i-1]{
            continue
        }
        newSlice := append([]int{},candidates[i+1:]...)
        convination = append([]int{candidate},convination[:]...)
        if solution(newSlice,(target-candidate),convination,res){
            *res = append(*res,convination)
            return false
        }
        convination = append([]int{},convination[1:]...)
    }
    return false
}