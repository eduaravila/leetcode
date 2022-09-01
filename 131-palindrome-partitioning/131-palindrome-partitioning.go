func isPalindrome(input string)bool{
    if len(input) < 1{
        return true
    }
    
    l,r := 0, len(input)-1
    
    for l < r {
        if input[l] != input[r]{
            return false
        }
        l++
        r--
    }
    
    return true 
}

func partition(s string) [][]string {
    res := [][]string{}
    solution(s,0,[]string{},&res)
    return res
}

func solution(s string,current int,subString []string, res *[][]string ){
    
    if current >= len(s){        
        *res = append(*res,append([]string{},subString...))
        return
    }
    
    for i := current ; i < len(s) ; i++ {
        if isPalindrome(s[current:i+1]){
            subString  = append(subString, s[current:i+1])
            
            solution(s,i+1,subString,res)
            subString = subString[:len(subString)-1]
        }
    }
    
}