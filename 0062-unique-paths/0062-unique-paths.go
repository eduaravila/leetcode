func swap(a,b *[]int){
    temp := *a    
    *a = *b
    *b = temp
}

func uniquePaths(m int, n int) int {
    col,row := []int{},[]int{}
    for i := 0; i < n; i++{
        col = append(col,1)
        row = append(row,1)
    }
    for r:= 1 ; r < m ;r++ {
        for c := 1 ; c < n ; c++ {
            col[c] = row[c] + col[c-1]
        }
        swap(&col,&row)
    }
    return row[n-1]
}