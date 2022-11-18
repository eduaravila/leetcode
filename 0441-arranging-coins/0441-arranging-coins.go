func arrangeCoins(n int) int {
    i := 1
    var res int
    for n>=0{
        
        n-=i
        if n >= 0{
            res++
        }
        
        i++
    }
    return res
}



// 1, 5
// 3
// 1 4
// 2 2
// 3 -1