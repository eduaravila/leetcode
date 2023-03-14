type key struct {
    a,b int
}

func combinationSum4(nums []int, target int) int {
    return solution(nums, target, 0, make(map[key]int))       
}

func solution(nums []int, target, indx int, memo map[key]int ) int{
    k := key{target,indx}
    if val, e := memo[k];e {
        return val
    }
    
    if target == 0{
        return 1
    }
    
    if indx > len(nums) -1 {
        return 0
    }
    
    if target < 0{
        return 0
    }
    
    var res int
    
    for i, num := range nums {
        res += solution(nums,target - num, i, memo)
    }
    memo[k] = res
    return res
}

/*

- different sequences are counted as different 


- base case 
 if target == 0{
     return 1
 }
 
 if target < 0 || i < 0 {
     return 0
 }
 
 
- have 2 different indexes 
- 1 for the current loop
- 2 for the recursion

- case 1:

var res int
for i < len(nums){

    res += solution(nums, i, target - nums[i])
    
    - increase i
    i++
    
}

1 
2 1 1
3



*/