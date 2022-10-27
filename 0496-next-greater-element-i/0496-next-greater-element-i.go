func nextGreaterElement(nums1 []int, nums2 []int) []int {
    indexes := make(map[int]int)
    
    for i, num := range nums2{
        indexes[num] = i
    }
    res := []int{}
    for _, num := range nums1{
        l := indexes[num]
        great := -1
        for l < len(nums2){
            if nums2[l] > num{
                break
            }
            l++
        }
        if l < len(nums2){
            great = nums2[l]   
        }
        res = append(res,great)    
    }
    return res
}