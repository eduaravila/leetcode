func strStr(haystack string, needle string) int {
    for l:=0 ; ; l++{
        for r:=0 ;; r++{
            if l+r > len(haystack)-1{return -1}
            if r > len(needle)-1 {break}
            if haystack[l+r] != needle[r]{break}
            if l+r - l +1 == len(needle){return l}
        }
    }
    return -1
}