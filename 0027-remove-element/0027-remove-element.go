func removeElement(nums []int, val int) int {
    total := 0
    i,r := 0,len(nums)-1
    for i < len(nums){
        if nums[i] == val {
            nums[i], nums[r] = nums[r], -1
            r--
            total++
        }else{
            i++
        }
    }
    return len(nums) - total
}