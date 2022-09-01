import "strings"

var letters = map[string][]string{
    "2":[]string{"a","b","c"},
    "3":[]string{"d","e","f"},
    "4":[]string{"g","h","i"},
    "5":[]string{"j","k","l"},
    "6":[]string{"m","n","o"},
    "7":[]string{"p","q","r","s"},
    "8":[]string{"t","u","v"},
    "9":[]string{"w","x","y","z"},
}
func letterCombinations(digits string) []string {
    res := []string{}
    solution(strings.Split(digits,""),&res,[]string{})
    return res
}

func solution(digits []string,res *[]string,current []string)bool{
    if len(digits) <1 {
        return true
    }
    
    for _,letter := range letters[digits[0]] {        
        current = append(current,letter)
        if solution(digits[1:],res,current){
            *res= append(*res,strings.Join(current,""))
        }
        current = append([]string{},current[:len(current)-1]...)
    } 
    
    return false
}