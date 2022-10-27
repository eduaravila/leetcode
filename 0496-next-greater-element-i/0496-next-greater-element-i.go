func nextGreaterElement(nums1 []int, nums2 []int) []int {
    indexes := make(map[int]int)
    
    for i, num := range nums2{
        indexes[num] = i
    }
    res := []int{}
    for _, num := range nums1{
        l := indexes[num]
        for l < len(nums2){
            if nums2[l] > num{
                break
            }
            l++
        }
        if l < len(nums2){
            res = append(res,nums2[l])    
        }else{
            res = append(res,-1)
        }
    }
    return res
}