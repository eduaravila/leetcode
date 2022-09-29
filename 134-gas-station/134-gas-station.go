func canCompleteCircuit(gas []int, cost []int) int {
    startingPoints := []int{}
    
    for i := range gas{
        if gas[i] >= cost[i]{
            startingPoints = append(startingPoints, i)
        }
    }
    for len(startingPoints) > 0{
        start := startingPoints[len(startingPoints)-1]
        startingPoints = startingPoints[:len(startingPoints)-1]
        current := (start + 1 )% len(gas)
        cGas := (gas[start] - cost[start]) + gas[current]
        
        for cGas >= cost[current]   {
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