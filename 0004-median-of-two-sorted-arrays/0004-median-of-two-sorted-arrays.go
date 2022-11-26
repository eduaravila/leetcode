func getMax(a,b int)int{
    if a > b{
        return a
    }
    return b
}

func getMin(a,b int)int{
    if a < b{
        return a
    }
    return b
}


func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    total := len(nums1) + len(nums2)
    half := (total / 2 )    
    if len(nums2) < len(nums1){
        nums1,nums2 = nums2,nums1
    }
    l,r := 0, len(nums1)-1
    inf := int(^uint(0)>>1)
    var aleft,aright,bleft,bright int
    
    for {
        sa := int((l+r)>>1)
        sb := getMax(half - sa - 2,-1)
        
        if sa >= 0 {
            aright = nums1[sa]
        }else{
            aright = -inf
        }
        
        if sb >= 0{
            bright = nums2[sb]
        }else{
            bright = -inf
        }
        
        if sa +1 < len(nums1){
            aleft = nums1[sa+1]
        }else{
            aleft = inf
        }
        
        if  sb +1 < len(nums2){
            bleft = nums2[sb+1]
        }else{
            bleft = inf
        }
        
        if aright <= bleft && bright <= aleft{  

            if total % 2 == 0{                
                return (float64(getMax(aright,bright)) + float64(getMin(aleft,bleft))) / 2
            }
            return float64(getMin(aleft,bleft))
        }else if aright > bleft{
            r = sa-1
        }else{
            l = sa+1
        }
        
    }
    
    return float64(0)
}


