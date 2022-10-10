func increasingTriplet(nums []int) bool {
    inf := int(^uint(0)>>1)
    l,m := inf,inf
    
    for _,num := range nums{
        if num <= l{
            l = num
        }else if num <= m{
            m = num
        }else{
            return true
        }
    }
    return false
}