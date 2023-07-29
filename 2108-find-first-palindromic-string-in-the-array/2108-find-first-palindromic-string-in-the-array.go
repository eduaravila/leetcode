func firstPalindrome(words []string) string {
    for _, word := range words {
        if isPalindrome(word){
            return word
        }
    }    
    
    return ""
}

func isPalindrome(input string)bool{
    a, b := 0 ,len(input)-1
    
    for a <= b {
        if input[a] != input[b]{
            return false
        }
        a++
        b--
    }   
    return true
}