
// Input: s = "(*)"
// stack = 1 0 -> true, 1 2 1 -> false, 1 0 -1
// index = 0 1 2 , 2 ,2



// * = ( " " ) tree possible cases

// return stack == 0 
// " " = ignore and continue
// "(" add to stack and continue
// ")" pop from stack and continue

//  ( 
//  1-0
// "" - ( - )
// 1-1 2-1 0-1
// ) ) )
// 0-2 1-2 -1-2

func checkValidString(s string) bool {
    return solution(s, 0,0, make(map[string]bool))
}

func solution(s string, current, stack int, memo map[string]bool)bool{
    key := fmt.Sprintf("%d-%d",current,stack)
    if val,e := memo[key];e{
        return val
    }
    if stack < 0 || current > len(s)-1 && stack > 0 {
        return false
    }
    if current > len(s)-1 && stack==0{
        return true
    }
    
    if s[current] == '('{
        memo[key] = solution(s,current+1,stack+1,memo)
    }else if s[current] == ')'{
        memo[key] = solution(s,current+1,stack-1,memo)
    }else{
        memo[key] = solution(s,current+1,stack+1,memo) || solution(s,current+1,stack-1,memo) || solution(s,current+1,stack,memo)
    }
    return memo[key]
}