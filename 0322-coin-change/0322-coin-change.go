func getMin(a,b int)int{
    if a < b{
        return a
    }
    return b
}


func coinChange(coins []int, amount int) int {
    table := []int{}
    for i := 0 ; i <= amount;  i++{
        table = append(table, amount+1)
    }
    
    table[0] = 0
    
    for _,coin := range coins {        
        for i := 0; i <= amount; i++{            
            if i - coin < 0 {
                continue
            }
            table[i] = getMin(table[i], 1+table[i-coin])
        }
    }
    
    if table[amount] == amount+1{
        return -1
    }
    return table[amount]
}


