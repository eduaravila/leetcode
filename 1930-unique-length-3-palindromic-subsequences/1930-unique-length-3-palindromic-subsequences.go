type key struct{
    middle rune
    end rune
}

func isPalindrome(pre map[rune]bool, post map[rune]int,c rune, res map[key]bool){
    
    for i := 0; i < 26; i++{
        char := rune('a'+i)        
        if pre[char] && post[char] > 0{  
            res[key{c,char}] = true            
        }
    }
    
}

func countPalindromicSubsequence(s string) int {
    pre := make(map[rune]bool)
    post := make(map[rune]int)
    res := make(map[key]bool)
    for _, c := range s{
        post[c]++
    }
    
    for _,c := range s{        
        post[c]--
        isPalindrome(pre,post,c,res)
        pre[c] = true                
    }
    return len(res)
}