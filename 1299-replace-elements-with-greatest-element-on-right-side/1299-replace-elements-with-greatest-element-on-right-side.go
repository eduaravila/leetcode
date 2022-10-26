func getMax(a,b int)int{
    if a > b {
        return a
    }
    return b
}
func replaceElements(arr []int) []int {
    max :=-1
    r := len(arr)-1
    var temp int
    for r >= 0 {
        temp =arr[r]
        arr[r] = max        
        max = getMax(max,temp)
        r--
    }
    return arr
}