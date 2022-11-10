func rearrangeArray(nums []int) []int {
    sort.Ints(nums)
    m := int(0+ len(nums)-1)>>1    
    res := make([]int,len(nums))
    odd:= 0
    i:=0
    for i <=m {
        res[odd] = nums[i]
        i++
        odd+=2
    }
    pair := 1
    i= m+1
    for pair < len(nums){
        res[pair] = nums[i]
        pair+=2    
        i++
    }
    
    return res
}