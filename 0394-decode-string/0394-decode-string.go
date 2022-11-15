func decodeString(s string) string {    
    return solution(s)    
}

func solution(s string) string{    
    times := []int{}
    var i int
    patterns := []strings.Builder{strings.Builder{}}
    
    for i < len(s){            
        
        if s[i] >= '0' && s[i] <='9'{
            index := strings.Index(s[i:],"[")
            k, _ := strconv.Atoi(s[i:index+i])            
            times = append(times, k)
            patterns = append(patterns, strings.Builder{})
            i += index
            continue
        }
        
        if s[i] == '['{
            i++
            continue   
        }
        
        if s[i] == ']'{
            k := times[len(times)-1]
            times = times[:len(times)-1]
            temp:= patterns[len(patterns)-1]
            patterns = patterns[:len(patterns)-1]            
            tempp := patterns[len(patterns)-1].String()
            patterns[len(patterns)-1].Reset()
            patterns[len(patterns)-1].WriteString(tempp+strings.Repeat(temp.String(),k))            
            i++
            continue   
        }
        patterns[len(patterns)-1].WriteString(string(s[i]))
        i++        
    }
    return patterns[0].String()
}
