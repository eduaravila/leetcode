func canCompleteCircuit(gas []int, cost []int) int {
    visited := make([]bool,len(gas))        
    for start := range gas{
        if visited[start] || gas[start] < cost[start]{
            continue
        }
                    
        current := (start + 1 )% len(gas)
        cGas := (gas[start] - cost[start]) + gas[current]
        
        for cGas >= cost[current] {
            
            visited[current] = true
            if current == start {
                return start
            }
            
            cGas =  (cGas - cost[current])      
            current = (current + 1 )% len(gas)
            cGas += gas[current]
            
        }
    }
    return -1 
}