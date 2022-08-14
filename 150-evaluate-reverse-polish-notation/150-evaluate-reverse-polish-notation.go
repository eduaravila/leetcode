import ("fmt")
func performOperation(a,b int, operator string) int{
    if operator == "+"{
        return a+b
    }
    if operator == "*"{
        return a*b
    }
    if operator == "/"{
        return a/b
    }
    if operator == "-"{
        return a-b
    }
    return 0
}

func evalRPN(tokens []string) int {
    nums_stack := []int{}
    
    for _, token := range tokens {
        if _,e := strconv.Atoi(token); e != nil {
            a,b :=nums_stack[len(nums_stack)-1],nums_stack[len(nums_stack)-2]                                   
            nums_stack = nums_stack[:len(nums_stack)-2]
            fmt.Println(a,b,token)
            nums_stack  = append(nums_stack,performOperation(b,a,token))            
        }else{
            var num int
            fmt.Sscanf(token,"%d",&num)
            nums_stack = append(nums_stack,num)
        }

    }
    
    return nums_stack[0]
}