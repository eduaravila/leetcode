func canPlaceFlowers(flowerbed []int, n int) bool {
    var current int
    
    for n > 0 {
        if current >= len(flowerbed) {
            return false
        }
        
        if flowerbed[current] == 1 {
            current++
            continue
        }else if current > 0 && flowerbed[current-1] == 1{
            current++
            continue
        }else if current < len(flowerbed)-1 && flowerbed[current +1] == 1{
            current++
            continue
        }
        n--
        flowerbed[current] = 1
        current++
    }
    return true
}