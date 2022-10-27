func getLocalName(email string)string{    
    res := []string{}
    for _,c := range email{
        if c == '.'{
            continue
        }
        if c == '+' || c== '@' {
            break
        }
        res = append(res,string(c))
    }
    return strings.Join(res,"")
}

func getDomainName(email string)string{
    inx := 0
    for i,c := range email{
        inx = i
        if c == '@'{
            break
        }
        
    }
    
    return strings.Join(strings.Split(email,"")[inx:],"")
}

func numUniqueEmails(emails []string) int {
    mapping := make(map[string]bool)
    for _, email := range emails{
        key := fmt.Sprintf("%s%s",getLocalName(email),getDomainName(email))
        mapping[key] =true
    }
    fmt.Println(mapping)
    return len(mapping)
}