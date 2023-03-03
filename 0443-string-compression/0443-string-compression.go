func digitsToText(counter int)[]byte{

    if counter < 2{
        return []byte{}
    }
    var res string    
    var digit int    
    for counter > 0{        
        digit = counter % 10
        counter /=  10
        i := strconv.Itoa(digit)
        res = i + res
    }
    return []byte(res)
}


func compress(chars []byte) int {
    
    l, r := 0, 0
    var counter int
    var current byte
    for r < len(chars) {
        c := chars[r]
        if c == current{
            counter++
        }else {            
            if counter > 0{
                fmt.Println(counter,c)
                currentCompression := append([]byte{current},digitsToText(counter)...)                
                chars = append(chars[:l], append(currentCompression, chars[r:]...)...)
                l += len(currentCompression)
                r = l
            }
            current = c
            counter = 1
        }        
        r++
    }
    
    currentCompression := append([]byte{current},digitsToText(counter)...)                
    chars = append(chars[:l], append(currentCompression, chars[r:]...)...)
    
    return len(chars)
}

/*
aabbc3
aab2c3

current 2 
i 3, 1
r 3, 1


*/