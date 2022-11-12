func generateParenthesis(n int) []string {
    res := []string{}
    solution(n,n,[]rune{},&res)
    return res
}

func solution(o,c int, current []rune, res *[]string){
    if c < o || o < 0{
        return 
    }
    if  o < 1 && c == o{
        *res = append(*res,string(current))
        return
    }
    
    
    solution(o-1,c, append(current,'('),res)    
    solution(o,c-1, append(current,')'),res)
}