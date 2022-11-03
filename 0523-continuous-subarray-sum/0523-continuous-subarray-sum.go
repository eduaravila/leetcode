func checkSubarraySum(nums []int, k int) bool {    
    presums :=  make(map[int]int)
    presums[0] = 0
    var sum, i int    
    for i < len(nums){                             
        sum += nums[i]  
        remaining := sum % k 
        
        if _,e := presums[remaining]; !e {            
            presums[remaining] = i + 1                 
        }else if presums[remaining] < i{
            return true
        }
        i++
    }
    
    return false
    
}

// [23,2,4,6,7]
// [5:1,1:2, ]