func dailyTemperatures(temperatures []int) []int {
    stack := [][]int{}//index,val
    res := make([]int,len(temperatures))
    for i, val := range temperatures{
        for len(stack) > 0 && val > stack[len(stack)-1][1]{ // current > peek 
            key := stack[len(stack)-1][0] 
            res[key] = i - key
            stack = stack[:len(stack)-1]
        }
        stack = append(stack,[]int{i,val})
    }
    return res
}