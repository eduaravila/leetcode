func subarraySum(nums []int, k int) int {
    var res, sum int
    
    
    for i := range nums{
        sum = 0
        for y := i ; y < len(nums); y++{
            sum += nums[y]
            if sum == k{
                res++
                
            }
            
        }
    }
    return res
}


func solution(nums []int, k,current int, res *int){
    
    if len(nums) < 1{
        return
    }
    
    solution(append([]int{},nums[1:]...),k,current+nums[0],res)
    if current == k{
        *res++
    }
    
    solution(append([]int{},nums[1:]...),k,nums[0],res)

}