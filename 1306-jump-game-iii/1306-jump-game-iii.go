func canReach(arr []int, start int) bool {
    if len(arr) < 1{
        return true
    }
    
    visited := make([]bool,len(arr))
    queue := []int{start}
    
    for len(queue) > 0{
        
        current := queue[len(queue)-1]
        queue = queue[:len(queue)-1]
        if visited[current] {
            continue
        }
        if current + arr[current] < len(arr){
            queue = append(queue,current + arr[current])
        }
        
        if current - arr[current] > -1  {
            queue = append(queue,current - arr[current])
        }
        
        visited[current] = true
        if arr[current] == 0 {
            return true
        }
    }
    
    
    return false
}