import (
    goheap "container/heap"
)

type heap []int

func (h *heap) Less(a,b int) bool{return (*h)[a] > (*h)[b] }
func (h *heap) Len() int { return len(*h)}
func (h *heap) Swap(a,b int) { (*h)[a],(*h)[b] = (*h)[b],(*h)[a] }
func (h *heap) Pop() interface{} {
    temp := (*h)[len(*h)-1]
    *h = (*h)[:len(*h)-1]
    return temp
}
func (h *heap) Push(val interface{}) {*h = append(*h,val.(int))}

func maxSlidingWindow(nums []int, k int) []int {
    _h := &heap{}
    goheap.Init(_h)
    
    for i:=0 ; i < k; i++{
        goheap.Push(_h, nums[i])        
    }    
    res := []int{}
    l := 0
    res = append(res,(*_h)[0])
    for r := k; r < len(nums); r++{
        goheap.Push(_h, nums[r])
        idx := 0 
        for i,val := range *_h{
            if val == nums[l]{
                idx= i
                break
            }
        }
        goheap.Remove(_h,idx)
       
        res = append(res, (*_h)[0])
        l++
    }
    return res
}