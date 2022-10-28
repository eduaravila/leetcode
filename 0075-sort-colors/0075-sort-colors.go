func sortColors(nums []int)  {
    
    sort.Slice(nums, func(a,b int)bool{
        return nums[a] < nums[b]
    })
    
}