func getMin(a,b int)int{
    if a < b{
        return a        
    }    
    return b
}

func suggestedProducts(products []string, searchWord string) [][]string {
    res := [][]string{}
    sort.Strings(products)
    l,r := 0,len(products)-1
    for i :=1; i <= len(searchWord);i++{        
        word :=  searchWord[:i]
        for l < len(products) && !strings.HasPrefix(products[l], word){
            l++
        }        
        for r >= 0 && !strings.HasPrefix(products[r], word){
            r--
        }
        
        if l < len(products) && r >= 0{
            res = append(res,products[l:l+getMin(3,r-l+1)])
        }else{
            res =append(res,[]string{})
        }
        
    }
    return res    
}