func productExceptSelf(nums []int) []int {
    left ,right := []int{1},[]int{1}
    i := 0
    for i < len(nums)-1{
        left = append(left, left[i] * nums[i])
        i++
    }
    
    i = len(nums)-1
    for i > 0{
        right = append(right, right[(len(nums)-1)-i] * nums[i])
        i--
    }
    res := []int{}
    for i := range left{
        res = append(res , left[i]* right[(len(right)-1)-i])
    }
    return res
}