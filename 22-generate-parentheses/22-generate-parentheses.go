func generateParenthesis(n int) []string {
    res := []string{}
    solution([]rune{},0,n,&res)
    return res
}

func solution(current []rune, stack, n int, res *[]string)bool{
    if n < 0{
        return true
    }
    
    r := solution(append(current,'('),stack+1,n-1,res)
    if stack > 0 {
        solution(append(current,')'),stack-1,n,res)
    }
    
    if stack <= 0 && r { // n and stack empty
        *res = append(*res,string(current))
    }
    
    return false
}