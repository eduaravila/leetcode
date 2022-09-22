func getMaxIndex(nums []int, l,r int)int{
    if nums[l] == 0{
        return l
    }
    inf := int(^uint(0)>>1)
    min := -inf
    maxI := l+1
    for i := l+1 ; i <= r ; i++{
        if i > len(nums) - 1 {
            break
        }
        if nums[i] >= min{
            min = nums[i]
            maxI = i
        }
    }
    return maxI
}

func canJump(nums []int) bool {
    if len(nums) <2 {
        return true
    }
    
    stack := []int{0}
    for len(stack) > 0{        
        i := stack[len(stack)-1]
        current := nums[i]        
        max := getMaxIndex(nums,i,current+i)
        if i == len(nums) -1 || max == len(nums)-1 {
            return true
        }
        if nums[max] == 0{
            nums[i] = 0
            stack = stack[:len(stack)-1]
        }else{            
            stack = append(stack,max)
        }
        
        
    }
    return false
}