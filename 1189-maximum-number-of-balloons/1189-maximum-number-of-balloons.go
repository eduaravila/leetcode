func maxNumberOfBalloons(text string) int {
    counter := make(map[rune]int)
    for _,c := range text{
        counter[c]++
    }
    
    var res int
    for {
        for _, c:= range "balloon"{
            if counter[c]> 0{
                counter[c]--
            }else{
                return res
            }
            
        }
        res++
    }

    return res
}