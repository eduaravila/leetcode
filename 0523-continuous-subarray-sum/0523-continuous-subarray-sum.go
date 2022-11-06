func checkSubarraySum(nums []int, k int) bool {
    pref := make(map[int]int)
    pref[0] = 0
    var sum int
    for i,num := range nums{        
        sum+=num        
        if _,e := pref[sum%k];!e{
            pref[sum%k] = i+1            
        }else if pref[sum%k] < i{
            return true
        }
    }
    return false
}