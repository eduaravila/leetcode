/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type Result struct{
    Node *ListNode
    Valid bool
}

func getSize(head *ListNode)int{
    var res int
    for head !=nil{
        res++
        head =head.Next
    }
    return res
}

func isPalindrome(head *ListNode) bool {
    size := getSize(head)
    return solution(head,size).Valid
}


func solution(head *ListNode, size int)Result{
    if size == 0 || size == 1{
        if size == 1 {
            return Result{head.Next,true}
        }
        return Result{head,true}
    }
    
    res := solution(head.Next,size-2)  
    if !res.Valid{
        return res
    }
    if res.Node.Val != head.Val{
       res.Valid = false
    }
    
    res.Node = res.Node.Next
    return res
    
}