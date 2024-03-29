func merge(nums1 []int, m int, nums2 []int, n int)  {
    
    for m > 0 && n > 0{
        if nums1[m-1] > nums2[n-1]{
            nums1[m+n-1] = nums1[m-1]
            m--
        }else {
            nums1[m+n-1] = nums2[n-1]
            n--
        }
    }
    
    for i := 0 ; i < n ; i++ {
        nums1[i] = nums2[i]        
    }
}