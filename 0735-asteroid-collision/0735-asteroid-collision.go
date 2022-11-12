func asteroidCollision(asteroids []int) []int {
    res := []int{}    
    
    for _,asteroid := range asteroids{
        if asteroid > 0{
            res = append(res, asteroid)
            continue   
        }
        
        ok := true
        for len(res) > 0  && res[len(res)-1] > 0{
            pos :=  res[len(res)-1]
            if int(math.Abs(float64(asteroid))) >= pos{
                res = res[:len(res)-1]                
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
    
    return res
    
}



// p 10 2
// n -5

