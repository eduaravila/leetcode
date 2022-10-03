func checkValidString(s string) bool {    
    return solution(0,0,strings.Split(s,""), make(map[string]bool))
}

func solution(size,current int, s []string, memo map[string]bool)bool{
    key := fmt.Sprintf("%d,%d",size, current)
       
    if size == len(s) || current < 0 {
        return current == 0
    } 
    if val,e := memo[key];e{
        return val
    } 
    if s[size] == "("{
        memo[key] = solution(size+1,current+1,s,memo)
    }else if s[size] == ")" {
        memo[key] = solution(size+1,current-1,s,memo)
    }else {                        
        memo[key] = solution(size+1, current+1,s,memo) || solution(size+1, current-1,s,memo) || solution(size+1,current,s,memo)
    }
    return memo[key]
}