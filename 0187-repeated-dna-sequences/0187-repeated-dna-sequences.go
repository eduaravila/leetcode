func findRepeatedDnaSequences(s string) []string {
    if len(s) < 10{
        return []string{}
    }
    var mask int
    s1 := make(map[int]int)
    s2 := make(map[int]bool)
    res := []string{}
    for i := 0 ; i< 10;i++{
        mask = mask << 2 | mapToBit(s[i])
    }
    s1[mask]++
    initial := (1 << 20) -1
    
    for i:= 10; i< len(s);i++{        
        mask = (mask << 2) & initial | mapToBit(s[i])  
        
        if s2[mask]{
            continue
        }

        if s1[mask] > 0{
            s2[mask]= true
            res = append(res,s[i-9:i+1])
        }
        s1[mask]++
    }
    return res
}

func mapToBit(s byte)int{
    switch(s){
        case 'A':
            return 0
        case 'C':
            return 1
        case 'G':
            return 2
        case 'T':
            return 3
    }
    return 0
}