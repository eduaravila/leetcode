func lengthOfLongestSubstring(s string) int {
    l,r ,res := 0,0,0
    subStringCounter := make(map[byte]bool)

    for r < len(s) {                     
        if _,e := subStringCounter[s[r]]; !e{            
            subStringCounter[s[r]] = true            
            r++
            if len(subStringCounter) > res{
                res = len(subStringCounter)
            }
        }else{
            delete(subStringCounter,s[l])
            l++
        }
        
    }
    return res
}

// start a window
// count the windows elements
// if we encounter a repeated element inside the window
// delete from left and shrink by increasing l++
// if the element is new add it to the substring and compare the result