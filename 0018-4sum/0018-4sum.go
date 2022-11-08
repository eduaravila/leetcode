func fourSum(nums []int, target int) [][]int {
    sort.Ints(nums)
    res := [][]int{}
    solution(nums,[]int{}, &res,0, target,4)
    return res
}

func solution(nums,current []int,res *[][]int, index, target, k int){
    
    if k == 2 {   
        
        l,r := index, len(nums)-1
        for l < r{
            sum := nums[l] + nums[r]
            if sum == target{
                *res = append(*res, append([]int{nums[l],nums[r]},current...))                
                r--
                for l < r && nums[r] == nums[r+1]{
                    r--
                }
            }else if sum > target {                
                r--
            }else{                
                l++
            }
        }
        return 
    }
    
    for i:= index ; i < len(nums) - k + 1; i++{        
        if i > index && nums[i] == nums[i-1]{ // omit duplicated starts            
            continue
        }
        current = append(current,nums[i])
        solution(nums, current, res, i+1, target-nums[i], k-1)
        current = append([]int{},current[:len(current)-1]...)
        
    }
}