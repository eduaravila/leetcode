func canPlaceFlowers(flowerbed []int, n int) bool {
    
    if len(flowerbed) < 2 && flowerbed[0] == 0{
        
        return n-1 < 1
    }
    if flowerbed[0] == 0 && flowerbed[1] == 0{
        n--
        flowerbed[0]=1
    }
    if flowerbed[len(flowerbed)-1] == 0 && flowerbed[len(flowerbed)-2] == 0{
        n--
        flowerbed[len(flowerbed)-1]=1
    }
    
    i := 1
    for i < len(flowerbed)-1 {        
        if flowerbed[i] == 0 && flowerbed[i-1] == 0 && flowerbed[i+1]==0 {
            n--
            flowerbed[i] = 1 
        }
        i++
    }
    
    
    return n < 1
}

