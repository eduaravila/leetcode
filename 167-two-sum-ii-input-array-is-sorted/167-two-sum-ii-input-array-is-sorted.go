func twoSum(numbers []int, target int) []int {
    l:=0
    r:=len(numbers)-1
    for l< r{                
        remaining := numbers[l] + numbers[r]
        
        if remaining < target {            
            l++            
        }else if remaining > target {
            r--
        }else{
            return []int{l+1,r+1}
        }
    }
    return []int{}
}