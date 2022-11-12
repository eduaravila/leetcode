func asteroidCollision(asteroids []int) []int {
    res := []int{}
    positives :=[]int{}
    
    for _,asteroid := range asteroids{
        if asteroid > 0{
            positives = append(positives, asteroid )
            continue   
        }
        
        ok := true
        for len(positives) > 0 {            
            pos :=  positives[len(positives)-1]
            if int(math.Abs(float64(asteroid))) >= pos{
                positives = positives[:len(positives)-1]                
            }
            
            if int(math.Abs(float64(asteroid))) <= pos{                
                ok = false
                break
            }                                     
        }
        if ok {
            res = append(res,asteroid)
        }
    }
    
    return append(res, positives...)
    
}



// p 10 2
// n -5

