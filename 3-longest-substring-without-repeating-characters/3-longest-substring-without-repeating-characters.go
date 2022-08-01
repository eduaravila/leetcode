func lengthOfLongestSubstring(s string) int {
    l,res := 0,0
    subStringCounter := make(map[rune]int)

    for r, run := range s{
        
        subStringCounter[run]++
        
        for l <= r && subStringCounter[run] > 1 {            
            subStringCounter[rune(s[l])]--
            l++
        }
        if (r - l) +1 > res{
            res= (r - l) + 1
        }
        
    }
    return res
}

// count temporal window elements
// get max substring
