func getMin(a,b int)int{
    if a < b{
        return a
    }
    return b
}


func minimumTotal(triangle [][]int) int {
    n := len(triangle)
    m := (len(triangle[n-1]))
    table := make([]int, m)
    
    for i, num := range triangle[n-1]{
        table[i] = num
    }
    
    for level := n-2 ; level >= 0;level--{
        for i := 0; i <= level; i++{
            table[i] = getMin(table[i], table[i+1]) + triangle[level][i]
        }
    }
    return table[0]
}

/*
    
       2
      3 4
     6 5 7
    4 1 8 3
    
    
    base case level == last level
        return triangle[level][i]
    
    table[i] = getMin(table[level+1][i], table[level+1][i+1]) + table[level][1]
    
*/

/*
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