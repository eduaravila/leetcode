func canCompleteCircuit(gas []int, cost []int) int {
    
    visited := make([]bool,len(gas))
    var current int
    for current < len(gas) {
        if gas[current] < cost[current] || visited[current] {
            current++
            continue
        }
        var tripSize int
        cGas, cStation := gas[current], current
        for tripSize < len(gas) {                        
            visited[cStation] = true
            nextStation := (cStation + 1)% len(gas)
            cGas = (cGas - cost[cStation]) 
            if cGas < 0 {
                break
            }
            cGas += gas[nextStation]
            cStation = nextStation            
            tripSize++
        }
        if tripSize == len(gas){
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