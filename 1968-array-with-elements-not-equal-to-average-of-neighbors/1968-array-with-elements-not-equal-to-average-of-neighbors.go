func rearrangeArray(nums []int) []int {
    sort.Ints(nums)
    m := (len(nums)-1) >> 1
    res := make([]int,len(nums))
    var pair int
    for i:=0; i <=m;i++{
        res[pair] = nums[i]
        pair+=2
    }
    
    odd := 1
    for i := m+1; i < len(nums);i++{
        res[odd] = nums[i]
        odd+=2
    }
    return res
}