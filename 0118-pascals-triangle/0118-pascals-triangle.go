func generate(numRows int) [][]int {
    res := [][]int{}
    var i int
    for i < numRows{
        row := []int{}
        for y:= 0; y<=i;y++{
            val := 1
            if i > 0 && y > 0 && y < i{
                val = res[i-1][y-1] + res[i-1][y]
            }
            row = append(row,val)
        }
        res =append(res, row)
        i++
    }
    return res
}