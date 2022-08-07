func characterReplacement(s string, k int) int {
    counter := make(map[byte]int)
    var max,res,l,r int             
    for r < len(s){
        
        counter[s[r]]++
        if counter[s[r]] > max {
            max = counter[s[r]]
        }
        if (r - l +1) - max >k {
            counter[s[l]]--
            l++
        } 
        if r - l +1 > res{
            res = r-l+1
        }
        r++
    }
    return res
}