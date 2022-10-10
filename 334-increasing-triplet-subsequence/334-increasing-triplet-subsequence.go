
func increasingTriplet(nums []int) bool {
    vals := []int{}
    for i := 0 ; i < 3 ; i++{
        vals = append(vals,int(^uint(0)>>1))
    }
    
    for _,num := range nums{
        var i int
        for i < len(vals){
            if vals[i] >= num{
                vals[i] = num
                break
            }
            i++
        }
        if i >=3 {
            return true
        }
    }
    return vals[2] != int(^uint(0)>>1)
}
