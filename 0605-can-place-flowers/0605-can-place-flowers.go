func canPlaceFlowers(flowerbed []int, n int) bool {
    var i int 
    for i < len(flowerbed) {        
        if flowerbed[i] == 0 && (i < 1 || (flowerbed[i] == 0 && flowerbed[i-1] == 0))  && (i >= len(flowerbed)-1 || (flowerbed[i] == 0 && flowerbed[i+1] == 0)) {
            n--
            flowerbed[i] = 1
        }
        i++
    }
    return n < 1
}

