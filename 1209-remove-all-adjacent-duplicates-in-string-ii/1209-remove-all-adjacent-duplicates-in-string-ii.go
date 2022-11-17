type key struct{
    c rune 
    count int
}

func removeDuplicates(s string, k int) string {
    sarr := strings.Split(s,"")
    stack := []key{}
    var i int
    for _, c:= range s{
        if len(stack) > 0 && stack[len(stack)-1].c == c{
            stack[len(stack)-1].count++
        }else{
            stack = append(stack,key{c,1})
        }
        if len(stack) > 0 && stack[len(stack)-1].count == k {
            stack = stack[:len(stack)-1]
            sarr = append(sarr[:i-k+1],sarr[i+1:]...)            
            i-=k
        }
        i++
    }
    return strings.Join(sarr,"")
}

//deeedbbcccbdaa
//ddbbbdaa
//d2b3
//0,4