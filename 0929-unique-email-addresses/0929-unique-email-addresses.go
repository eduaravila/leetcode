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
    for i,c := range email{        
        if c == '@'{
            return strings.Join(strings.Split(email,"")[i:],"")
        }     
    }
    
    return ""
}

func numUniqueEmails(emails []string) int {
    mapping := make(map[string]bool)
    for _, email := range emails{
        key := fmt.Sprintf("%s%s",getLocalName(email),getDomainName(email))
        mapping[key] =true
    }
    return len(mapping)
}