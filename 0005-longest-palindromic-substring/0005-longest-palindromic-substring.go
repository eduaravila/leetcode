

func longestPalindrome(s string) string {
    var max string
    
    for i := range s{
        solution(s,i,i,&max)
        solution(s,i,i+1,&max)
    }
    return max
}


func solution(s string, l,r int, max *string){
    
    var temp string
    for l > -1 && r < len(s){
        if s[l] != s[r]{
            break
        }        
        
        temp = s[l:r+1]
        l--
        r++
    }
    
    if len(temp) > len(*max){
        *max = temp
    }
}