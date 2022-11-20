func searchMatrix(matrix [][]int, target int) bool {
    
    l,r := 0,len(matrix)-1
    n:= len(matrix[0])-1
    var m int
    for l <= r{
        m = (l+r)>>1
        if target > matrix[m][n]{
            l = m+1
        }else if target < matrix[m][0]{
            r =m-1
        }else{
            break
        }
    }
    
    
    l,r = 0,len(matrix[m])-1
    for l <= r{
        mr:= (l+r)>>1
        if matrix[m][mr] == target{
            return true
        }
        if matrix[m][mr] < target{
            l=mr+1
        }else{
            r=mr-1
        }
    }
    return false
    
}

