type NumMatrix struct {
    pref [][]int
}


func Constructor(matrix [][]int) NumMatrix {
    pref := [][]int{}
    for x,row := range matrix{
        var sum, top int
        pref = append(pref,[]int{})
        for y,col := range row{            
            if x > 0{
                top = pref[x-1][y]
            }
            sum += col
            pref[x] = append(pref[x],sum+top)
        }
    }
    return NumMatrix{pref}
}


func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
    total := this.pref[row2][col2]
    if col1 > 0 {
        total -= this.pref[row2][col1-1]
    }
    if row1 > 0 {
        total -= this.pref[row1-1][col2]
    }
    if row1 > 0 && col1 > 0{
        total += this.pref[row1-1][col1-1]
    }
    return total
}


/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */