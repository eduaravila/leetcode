func getMin(a,b int)int{
    if a < b {
        return a
    }
    return b
}

func find132pattern(nums []int) bool {
    min := nums[0]
    stack := [][]int{}
    
    for _, num := range nums[1:] { 
        for len(stack) > 0 && num >= stack[len(stack)-1][0]{            
            stack = stack[:len(stack)-1]
        }
        if len(stack) > 0 && num > stack[len(stack)-1][1]{
            return true
        }  
        stack = append(stack,[]int{num,min})
        min = getMin(min, num)

    }    
    return false
}

//[3,5,0,3,4]
//5,0
//3,0

// 3,5,4
// 