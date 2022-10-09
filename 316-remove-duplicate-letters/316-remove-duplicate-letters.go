import "strings"
// count values 
// range(0,len(s))
// check if stack.peek() > current
	// do we have this value in the stack ? then we can ignore, we already have this value in place in the stack
	// do we have this value available in s[current+1:] ? then we can remove it, is bigger and we are sure we can replace it in future iterations

func removeDuplicateLetters(s string) string {    
    counter := make([]rune, ('z' - 'a')+1)
    for _,c := range s{
        counter[c-'a']++
    }
    
    var r int
    stack := []rune{}
    unique := make([]rune,('z' - 'a')+1)        

    for r < len(s){        
        for len(stack) > 0 && stack[len(stack)-1] > rune(s[r]) && counter[rune(stack[len(stack)-1])-'a'] > 0 && unique[rune(s[r])-'a'] < 1{                     
                unique[rune(stack[len(stack)-1])-'a']--
                stack = stack[:len(stack)-1]
        }        
        
        if unique[rune(s[r])-'a']<1{
            stack = append(stack,rune(s[r]))
            unique[rune(s[r])-'a']++
        }
        counter[rune(s[r])-'a']--
        r++
    }
    
    
    
    return string(stack)
}