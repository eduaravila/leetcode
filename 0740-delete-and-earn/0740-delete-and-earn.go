
func getMax(a,b int) int{
    if a > b{
        return a
    }
    return b
}

func deleteAndEarn(nums []int) int {
    counter := make(map[int]int)
    
    set := []int{}
    
    for _, num := range nums {
        if _, e := counter[num]; !e {            
            set = append(set, num)
        }
        counter[num]++        
    }
    sort.Ints(set)
    dp := make([]int, len(set))
    dp[0] = counter[set[0]] * set[0]
    
    i := 1
    for i < len(set) {
        num := set[i]
        current := counter[num] * num
        
        if i > 1 && set[i-1] == num-1 {
            dp[i] = getMax(dp[i-1], dp[i-2] +current)
        }else if i >1 {
            dp[i]= getMax(current + dp[i-1], current + dp[i-2])
        }else if set[i-1] < num-1{
            dp[i]= getMax(current, current + dp[i-1])
        }else{
            dp[i]= getMax(current, dp[i-1])
        }
        i++
        
    }
    fmt.Println(dp,set) 
    return dp[len(set)-1]
}

/*
[8,10,4,9,1,3,5,9,4,10]

[1,3,4,5,8,9,10]

*/