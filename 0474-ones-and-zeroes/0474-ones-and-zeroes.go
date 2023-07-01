func getMax(a,b int)int{
    if a> b{
        return a
    }
    return b
}

func findMaxFormsDp(pairs [][]int,current, availableZ, availableO int, memo map[string]int )int{
    key := fmt.Sprintf("%i,%i,%i",current,availableZ,availableO)
    if val, e := memo[key]; e{
        return val
    }
    if current >= len(pairs) || availableZ < 0 || availableO < 0 {
        return 0
    }
    
    if pairs[current][0] > availableZ  || pairs[current][1] > availableO {
        memo[key] = findMaxFormsDp(pairs,current+1,availableZ,availableO, memo)
        return memo[key]
    }
    
    include := 1 + findMaxFormsDp(pairs, current+1, availableZ - pairs[current][0], availableO - pairs[current][1], memo)
    exclude := findMaxFormsDp(pairs, current+1, availableZ, availableO, memo)
    memo[key] =  getMax(include, exclude)
    return memo[key]
}

func findMaxForm(strs []string, m int, n int) int {
    dp := [][]int{}
    sets := [][]int{}
        
    for i := 0 ; i <=m ; i++{
        dp = append(dp, make([]int,n+1))    
    }
    
    for _,str := range strs {
        ones,zeros := 0,0
        
        for _, c := range str{
            if c == '1'{
                ones++
            }else{
                zeros++
            }
        }
        sets = append(sets, []int{zeros,ones})
    }
    
    for _, set := range sets{
        for i := m ; i >= set[0]; i--{
            for x := n ; x>= set[1]; x--{
                dp[i][x] = getMax(dp[i][x], 1+ dp[i-set[0]][x-set[1]])
            }
        }
    }
    
    return dp[m][n]
}

/*
m = 0's 
n = 1's


base case 

m = 0
n = 0 

res = 0


["10","0001","111001","1","0"]

set = (zero, one)

[[1,1],[3,1],[2,3],[0,1],[1,0]]

[0,1],[1,0],[1,1],[3,1],[2,3]

[0,1],[1,1],[2,2],[5,3],[7,6]


["10","0","1"]

[1,0],[0,1],[1,1]

[1,0],[1,1],[2,2]


count the number on sets 

sort the sets


prefix sum
compare a pair with n and m until it satisfies the condition and return the index

dp[i][j] = max(dp[i][j],1+dp[i-currZeros][j-currOnes]);


[0,1],[1,0],[1,1],[3,1],[2,3],[0,1],[0,1],[0,2]

i = 0
x = 3
res = 2


[0, 2, 2, 2]



*/

/*
 sets := [][]int{}
    
    for _,str := range strs {
        ones,zeros := 0,0
        for _, c := range str{
            if c == '1'{
                ones++
            }else{
                zeros++
            }
        }
        sets = append(sets, []int{zeros,ones})
    }
    
    sort.Slice(sets, func(a,b int)bool{
        return sets[a][0] < sets[b][0] && sets[a][1] < sets[b][1]
    })
    

    fmt.Println(sets)
    var max int
    for i := range sets{   
        var prevZ,prevO int
        for currentSet := i ; currentSet< len(sets);currentSet++{
            set := sets[currentSet]
            prevZ += set[0]
            prevO += set[1]
            
            if prevZ > m || prevO > n{
                break
            }
            max = getMax(max, (currentSet-i)+1)        
        }
    }    
    return max
*/