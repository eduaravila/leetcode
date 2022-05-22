import 
(
    "strings"
)

func isSpecial(letter byte) bool{
    return letter < 'a' || letter > 'z'
}

func isLetter(letter rune) bool{
    return letter>='a' && letter <='z'
}

func isNumber(letter rune)bool{ 
    return letter >= '0' && letter <='9'
}

func isPalindrome(s string) bool {
    
    s = strings.ToLower(s)
    cleared_s := []rune{}
    for _, letter := range s {
        if isNumber(letter) || isLetter(letter){
            
            cleared_s = append(cleared_s,letter)
        }
    }
    l := 0
    r := len(cleared_s)-1
    fmt.Println(cleared_s)
    for l < r {                
        if cleared_s[l] != cleared_s[r]{
            return false
        }
        l++
        r--
    }
    return true
}