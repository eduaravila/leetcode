func moveZeroes(nums []int)  {
    l,r := 0,1
    
    for l < len(nums) && r < len(nums){
        for l < len(nums) && nums[l] != 0 {
            l++   
        }
        r = l
        for r < len(nums) && nums[r] == 0 {
            r++
        }
        if l < len(nums) && r < len(nums){
            nums[l],nums[r] = nums[r],nums[l]
        }
        
    }
}