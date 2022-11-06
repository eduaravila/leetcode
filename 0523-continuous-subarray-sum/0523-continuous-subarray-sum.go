func checkSubarraySum(nums []int, k int) bool {
    pref := make(map[int]int)
    pref[0] = 0
    var sum int
    for i,num := range nums{        
        sum+=num        
        if val,e := pref[sum%k];e{
            if val < i{
                return true
            }
        }else{
            pref[sum%k] = i+1
        }
    }
    return false
}