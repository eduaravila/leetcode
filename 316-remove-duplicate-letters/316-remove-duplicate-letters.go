import "strings"

func removeDuplicateLetters(s string) string {    
    counter := make([]rune, ('z' - 'a')+1)
    for _,c := range s{
        counter[c-'a']++
    }
    
    var r int
    stack := []rune{}
    res := make([]rune,('z' - 'a')+1)        

    for r < len(s){        
        for len(stack) > 0 && stack[len(stack)-1] > rune(s[r]) {
            
            if counter[rune(stack[len(stack)-1])-'a'] > 0 && res[rune(s[r])-'a'] < 1{
                res[rune(stack[len(stack)-1])-'a']--
                stack = stack[:len(stack)-1]            
            }else{
                break
            }
        }        
        
        if res[rune(s[r])-'a']<1{
            stack = append(stack,rune(s[r]))
            res[rune(s[r])-'a']++
        }
        counter[rune(s[r])-'a']--
        r++
    }
    
    
    
    return string(stack)
}