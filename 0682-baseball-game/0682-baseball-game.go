func calPoints(operations []string) int {
    records := []int{}
    var n int
    for _, operation := range operations{
        var res int    
        n = len(records)-1
        if operation == "+"{            
            res = records[n-1] + records[n]            
        }else if operation == "D"{
            res =  records[n] * 2
            
        }else if operation == "C"{
            records = records[:n]
            continue
        }else{
            res,_ = strconv.Atoi(operation)
        }
        
        records = append(records, res)
    }
    
    var res int
    
    for _, n := range records {
        res += n
    }
    
    return res
}

// 5, 10, 15

// 