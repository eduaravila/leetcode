/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type headWSize struct{
    Node *ListNode
    size int
}

func getSize(head *ListNode) int{
    var res int    
    current := head
    for current != nil{
        res++
        current = current.Next
    }
    return res
}

func getMin(a,b headWSize)headWSize{
    if a.size < b.size{
        return a
    }
    return b
}

func getMax(a,b headWSize)headWSize{
    if a.size > b.size{
        return a
    }
    return b
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
    hA,hB := headWSize{headA,getSize(headA)},headWSize{headB,getSize(headB)}
    currentMin, currentMax := hA, hB
    if hA.size != hB.size{
        currentMin =  getMin(hA,hB)
        currentMax = getMax(hA,hB)
    }
    
    for currentMin.size < currentMax.size{
        currentMax.size--
        currentMax.Node = currentMax.Node.Next
    }
    
    for currentMin.Node !=nil && currentMax.Node != nil{
        if currentMin.Node == currentMax.Node {
            return currentMin.Node
        }
        currentMax.Node = currentMax.Node.Next
        currentMin.Node = currentMin.Node.Next
    }
    return nil    
}