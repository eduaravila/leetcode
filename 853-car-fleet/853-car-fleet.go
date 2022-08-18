func getMin(a,b int)int {
    if a < b {
        return a
    }
    return b
}
func carFleet(target int, position []int, speed []int) int {
    convineTargetSpeed := [][]int{}
    for i,pos := range position {
        convineTargetSpeed = append(convineTargetSpeed,[]int{pos,speed[i]})
    }
    sort.Slice(convineTargetSpeed,func(a,b int)bool{
        return convineTargetSpeed[a][0] >= convineTargetSpeed[b][0]
    })
    
    fmt.Println(convineTargetSpeed)
    stack :=[]float64{}
    for _, val := range convineTargetSpeed{
        
        stack = append(stack, float64(target - val[0]) /float64(val[1]))
        
        if len(stack) >= 2 && stack[len(stack)-1] <= stack[len(stack)-2] {
            stack = stack[:len(stack)-1]
        }
        
    }
    return len(stack)
} 