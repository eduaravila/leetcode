func findTargetSumWays(nums []int, target int) int {
    return solution(nums,0,0,target)    
}

func solution(nums []int, i, sum ,target int)int{
    
    if i > len(nums)-1 && sum == target {
        return 1
    }
    if i > len(nums)-1 {
        return 0
    }
        
    return solution(nums, i+1, sum + (-nums[i]),target) +  solution(nums, i+1, sum + (+nums[i]),target)
}