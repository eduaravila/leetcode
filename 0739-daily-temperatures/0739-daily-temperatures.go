func dailyTemperatures(temperatures []int) []int {
    days := make([]int,len(temperatures))
    stack := []int{}
    for i := range temperatures{
        for len(stack)>0 && temperatures[i] > temperatures[stack[len(stack)-1]]{
            days[stack[len(stack)-1]] = i - stack[len(stack)-1]
            stack = stack[:len(stack)-1]
        }
        stack = append(stack,i)
    }
    return days
}

// 1, 2, 3, 4, 5
//  76
// 73, 74, 75
// 1, 1, 4, 2, 1,1,