func strStr(haystack string, needle string) int {
    
    for l:= 0; ; l++  {
        for r:=0 ; ; r++{
            if (r + l) - l == len(needle){return l}
            if r + l > len(haystack)-1 {return -1}
            if haystack[r+l] != needle[r]{break}
        } 
    }
    
    return -1
}