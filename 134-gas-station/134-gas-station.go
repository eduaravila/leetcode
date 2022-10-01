

func canCompleteCircuit(gas []int, cost []int) int {    
    
    var a,b int
    for i := range gas{
        a+= gas[i]
        b+= cost[i]
    }
    if a < b {
        return -1
    }
    visited := make([]bool,len(gas))
    var current int
    for current < len(gas){
        
        sum := gas[current] - cost[current] 
        if sum < 0 || visited[current]{
            current++
            continue
        }
        start := current
        for sum > 0 && current < len(gas)-1 {            
            current++
            sum += gas[current] - cost[current]
            
        }
        if current >= len(gas)-1{
            return start
        }
        current++
    }
    return -1
}

/*
tripSize 0 1 2 3 4 5
    cGas 4 8 7 6 5 0
cStation 3 4 0 1 2 3
current  3 
*/