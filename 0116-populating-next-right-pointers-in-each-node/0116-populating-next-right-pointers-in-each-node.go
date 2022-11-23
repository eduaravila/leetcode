/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
    if root==nil{
        return root
    }
    
    queue := []*Node{root}
    
    for len(queue) >0 {
        n := len(queue)
        for i, element := range queue{
            
            if element.Left!=nil{
                queue = append(queue,element.Left)
            }
            if element.Right!=nil{
                queue = append(queue,element.Right)    
            }

            if i+1 > n- 1{
                break
            }
            element.Next = queue[i+1]
            
        }
        queue = queue[n:]
    }
    
    return root
}