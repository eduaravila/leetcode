/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
    if node == nil{
        return nil
    }
    references := make(map[int]*Node)
     
    queue := []*Node{node}
    for len(queue) > 0{
        current := queue[0]
        queue = queue[1:]
        var clone *Node
        
        if val, e := references[current.Val]; e{
            clone = val
        }else{
            clone = &Node{Val: current.Val, Neighbors: []*Node{}}
            references[current.Val] = clone
        }
        
        for _, neighbor := range current.Neighbors{
            var neighborReference *Node
            if val, e:= references[neighbor.Val]; e{
                neighborReference = val
            }else{
                neighborReference = &Node{Val: neighbor.Val}
                references[neighbor.Val] = neighborReference
                queue = append(queue, neighbor)
            }
            
            clone.Neighbors = append(clone.Neighbors, neighborReference)
        }
        
    }
    
    return references[node.Val]
}