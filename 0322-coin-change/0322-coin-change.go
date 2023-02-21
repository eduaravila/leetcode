func getMin(a,b int)int{
    if a < b{
        return a
    }
    return b
}

const inf = int(^uint(0)>>1)
func coinChange(coins []int, amount int) int {
    table := []int{}
    for i := 0 ; i <= amount;  i++{
        table = append(table, inf)
    }
    
    table[0] = 0
    
    for _,coin := range coins {        
        for i := 0; i < amount; i++{            
            sum := coin + i            
            if table[i] == inf{
                continue
            }
            if sum > amount {
                break
            }
            
            ncoins := table[i] + 1
            table[sum] = getMin(table[sum], ncoins)
        }
    }
    
    if table[amount] == inf{
        return -1
    }
    fmt.Println(table)
    return table[amount]
}


