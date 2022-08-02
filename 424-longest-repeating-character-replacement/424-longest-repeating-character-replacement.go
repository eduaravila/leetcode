func getMax(a,b int) int{if a > b{return a}else{return b}}

func characterReplacement(s string, k int) int {
    l,r,res := 0,0,0
    counter := make(map[byte]int)
    maxSubString := 0
    for r< len(s){
        counter[s[r]]++
        maxSubString = getMax(maxSubString,counter[s[r]])
        
        if (r - l + 1) - maxSubString > k {
            counter[s[l]]--
            l++
        }
        res = getMax(res, r - l +1)
        r++
    }
    return res
}

// count the current r char occurrences
// check the r-l < k meaning we have n distance between but we can only replace 
// this will help us move l when we dont have any more replacements available
