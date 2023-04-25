func calculate(s string) int {
    sign, number, result := 1,0,0
    stack := []int{}
    
    for _,c := range s {
        
        if c == ' '{
            continue
        }
        if c >= '0' && c <= '9' {
            n,_ := strconv.Atoi(string(c))
            number = (10 * number  )+ n
        }else if c == '+'{
            result += number * sign
            sign = 1
            number = 0
        }else if c == '-'{           
            result += number * sign
            sign = -1 
            number = 0
        }else if c == '('{
            stack = append(stack, result)
            stack = append(stack, sign)
            sign =1
            result = 0
        }else if c == ')'{
            n := len(stack)
            result += sign * number
            prevs, prevn := stack[n-1], stack[n-2]
            number = 0
            stack = stack[:n-2]
            result *= prevs
            result += prevn
        }
    }
    if number !=0 {
        result += number * sign
    }

    return result    
}


/*
"(1+(4+5+2)-3)+(6+8)"



r = 4
s = 1
stack = 0,1,1,1


if c == num {
    number = 10 * number  + (number)
}

if c == +{
    sign = 1
    result = number * sing
    number = 0
}

if c == '('{
    stack = append(stack, result)
    stack = append(stack, sign)
    
    sign =1
    result = 0
}

if c == ')'{
    n := len(stack)
    prevs, prevn := stack[n-1], stack[n-2]
    
    result *= prevs
    result += prevn
}

if number !=0 {
    result += number * sign
}

return result
*/