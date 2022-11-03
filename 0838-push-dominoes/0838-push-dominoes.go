func pushDominoes(dominoes string) string {
    queue := []int{}   
    res := strings.Split(dominoes,"")

    for i,domino := range res{
        if domino != "."{
            queue=append(queue,i)
        }
    }    
    
    for len(queue) > 0{
        element := queue[0]
        queue = queue[1:]

        if len(queue) >0 && res[element] == "R" && queue[0]> element &&res[queue[0]]=="L" && queue[0] - (element+1) < 2 {            
            queue = queue[1:]
            continue
        }
        if res[element] == "L" && element-1 >=0 && res[element-1] =="."{
            res[element-1]="L"
            queue=append(queue,element-1)
        }else if res[element] == "R" && element+1 < len(dominoes) && res[element+1]=="." {
            res[element+1]="R"
            queue=append(queue,element+1)
        }
    }
    return strings.Join(res,"")
}