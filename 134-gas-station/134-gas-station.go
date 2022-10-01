

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
        
        sum, size,next := gas[current] - cost[current], 0,current
        if sum < 0 || visited[current]{
            current++
            continue
        }
        for sum > 0 && size < len(gas) {
            visited[next] = true
            next = (next +1 )%len(gas)
            sum += gas[next] - cost[next]
            size++
        }
        if size >= len(gas)-1{
            return current
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