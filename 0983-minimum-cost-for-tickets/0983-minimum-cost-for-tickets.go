func getMin(vals ...int)int{    
    if len(vals) < 1{
        return 0
    }
    min := vals[0]
    for _, val := range vals{        
        if val < min {
            min = val
        }
    }
    return min
}

func mincostTickets(days []int, costs []int) int {
    inf := int(^uint(0)>>1)
    return solution(days, costs, 0, len(days)-1, inf, make(map[key]int))
}

type key struct {
    a, b int
}

func solution(days, costs []int, sum, i, minDays int, memo map[key]int) int {    
    k := key{sum, minDays}
    
    if val, e := memo[k];e {
        return val
    }
    
    if i < 0 {        
        return sum
    }    
    
    if days[i] >= minDays {        
        res:= solution(days, costs, sum, i - 1, minDays, memo)
        memo[k] = res
        return res
    }
    res :=  getMin(
                solution(days, costs, sum + costs[0], i - 1, days[i], memo),
                solution(days, costs, sum + costs[1], i - 1, days[i] - 6, memo),
                solution(days, costs, sum + costs[2], i - 1, days[i] - 29, memo))    
    memo[k] = res
    return res
}

/*

base case 

if i < 0 {
    return sum
    // getMin here ? 
}

min = after doing current day - available number of days of the ticket 
this number tells us if is safe to continue to travel with the current ticket 

// still can travel with current ticket
if days[i] >= min{
    solution(days, cost, sum, i - 1, min)
}



solution(days, cost, sum + cost[0], i - 1, days[i] - 1)
solution(days, cost, sum + cost[1], i - 1, days[i] - 7)
solution(days, cost, sum + cost[2], i - 1, days[i] - 30)

*/