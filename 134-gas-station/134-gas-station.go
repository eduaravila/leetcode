func canCompleteCircuit(gas []int, cost []int) int {
    var tgas,tcost int
    
    for i := range gas{
        tgas+=gas[i]
        tcost+=cost[i]
    }
    if tcost > tgas{
        return -1
    }
    
    var i int
    for i < len(gas){
        if gas[i] < cost[i]{
            i++
            continue
        }
        sum := gas[i] - cost[i]
        r := i+1
        for sum >=0 && r < len(gas){
            sum += gas[r] - cost[r]
            r++
        }
        if r >=len(gas) || sum >=0{
            return i
        }
        i=r
    }
    return -1
}