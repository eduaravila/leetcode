func generateParenthesis(n int) []string {
    res := []string{}
    solution(n,n,[]string{},&res)
    return res
}

func solution(n int,s int, current []string,res *[]string){

    if n < s || n < 0 || s <0{
        return 
    }
    if n == 0 && s == 0{
        *res = append(*res,strings.Join(current,""))
        return 
    }        
    
    solution(n-1,s,append(current,")"),res)
    solution(n,s-1,append(current,"("),res)
    
}