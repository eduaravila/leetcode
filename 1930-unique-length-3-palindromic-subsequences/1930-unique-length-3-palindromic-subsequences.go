type key struct{
    mid rune
    end rune
}

func isPalindrome(middle rune, preffix map[rune]bool,suffix map[rune]int,res map[key]bool){
    
    for i :=0 ; i <26 ;i++{        
        letter := rune('a' + i)
        
        ep := preffix[letter]
        es := suffix[letter]
        if ep && es > 0 {
            res[key{middle,letter}]=true
        }
    }
    
}

func countPalindromicSubsequence(s string) int {
    preffix := make(map[rune]bool)
    suffix := make(map[rune]int)
    
    res := make(map[key]bool)
    for _,c :=range s{
        suffix[c]++
    }
    
    for _,c := range s{
        suffix[c]--
        isPalindrome(c,preffix,suffix,res)
        preffix[c] = true
        
    }
    
    
    return len(res)
}