func combinationSum2(candidates []int, target int) [][]int {
    res :=[][]int{}
    sort.Ints(candidates)
    solution(candidates,target,[]int{},&res)
    return res
} 

func solution(candidates []int, target int, current []int,res *[][]int)bool{
    if target == 0 {
        return true
    }
   
    
    for i,candidate := range candidates {
        if i > 0 && candidate == candidates[i-1]{
            continue
        }
        if candidate > target{
            break
        }
        if solution(append([]int{},candidates[i+1:]...),target-candidate,append(append([]int{},current...),candidate),res){
            *res = append(*res,append(current,candidate))
        }
    }
    return false
}