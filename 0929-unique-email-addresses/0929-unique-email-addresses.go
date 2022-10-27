func getLocalName(email string)string{    
    res := []rune{}
    for _,c := range email{
        if c == '.'{
            continue
        }
        if c == '+' || c== '@' {
            break
        }
        res = append(res,c)
    }
    return string(res)
}

func getDomainName(email string)string{    
    for i,c := range email{        
        if c == '@'{
            return email[i:]
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