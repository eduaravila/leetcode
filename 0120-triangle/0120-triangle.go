func getMin(a,b int)int{
    if a < b{
        return a
    }
    return b
}

func minimumTotal(triangle [][]int) int {
    return topDown(triangle, 0,0,make(map[string]int))
}

func topDown(triangle [][]int, i,level int, memo map[string]int)int{
    key := fmt.Sprintf("%d-%d",level, i)

    if val, e := memo[key];e{
        return val
        
    }
    
    if level == len(triangle)-1{
        return triangle[level][i]
    }
    
    memo[key] = getMin(topDown(triangle,i,level+1,memo),topDown(triangle,i+1,level+1,memo)) + triangle[level][i]
    return memo[key]
}


/*

  2
  3 4
 6 5 7
4 1 8 3



   11
  9 10
 7 6 10
4 1 8 3

*/