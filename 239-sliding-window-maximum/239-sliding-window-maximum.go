type heap []int

func NewHeap() heap{
    return []int{0}
}

func (h *heap) SUp(idx int)  {
    if idx < 1 || idx >= len(*h) {
        return
    }
    parent := idx / 2 
    
    if parent > 0 && (*h)[parent] < (*h)[idx] {
        h.swap(parent,idx)
        h.SUp(parent)  
    }
    
}

func (h *heap) SDown(idx int) {
    left, right := idx * 2, (idx * 2) + 1
    if left >= len(*h){
        return 
    }
    if right >= len(*h){
        right =left
    }
    if (*h)[left] > (*h)[idx] && (*h)[left] > (*h)[right] {
        h.swap(left,idx)
        h.SDown(left)
        return 
    }
    if (*h)[right] > (*h)[idx]{
        h.swap(right,idx)
        h.SDown(right)
        return
    }
    
}

func (h *heap) swap(a,b int){
     (*h)[a],(*h)[b] = (*h)[b], (*h)[a]
}

func (h *heap) Add(val int) *heap{
    *h = append(*h, val)
    h.SUp(len(*h)-1)
    return h
}

func (h *heap) Max() int{
    return (*h)[1]
}

func (h *heap) Remove(val int){
    idx := -1 
    for i:= 1 ;i < len(*h) ; i++ { 
        if (*h)[i] == val {
            idx = i
            break
        }
    }
    h.swap(idx,len(*h)-1)
    (*h) = (*h)[:len(*h)-1]        
    h.SDown(idx)
    h.SUp(idx)
}

func maxSlidingWindow(nums []int, k int) []int {
    h := NewHeap()
    for i:=0 ; i < k; i++{
        h.Add(nums[i])        
    }    
    res := []int{}
    l := 0
    res = append(res,h.Max())
    for r := k; r < len(nums); r++{
        h.Add(nums[r]) 
        h.Remove(nums[l])
       
        res = append(res, h.Max())
        l++
    }
    return res
}