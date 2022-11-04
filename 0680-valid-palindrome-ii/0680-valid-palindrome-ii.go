func isPalindrome(input string)bool{
l,r := 0,len(input)-1    
    for l < r {        
        if input[l] != input[r]{
              return false                 
        }
        l++
        r--
    }
    return true 
}

func validPalindrome(s string) bool {
    l,r := 0,len(s)-1    
    for l < r {        
        if s[l] != s[r]{
            nl := s[l:r]
            nr := s[l+1:r+1]
            return isPalindrome(nl) || isPalindrome(nr)            
        }
        l++
        r--
    }
    return true    
}



// abac