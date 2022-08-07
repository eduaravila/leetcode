

func maxSlidingWindow(nums []int, k int) []int {
    res := []int{}
    indexQ := []int{}
    valQ := []int{}
    for r,val := range nums{
        for len(valQ) > 0 && val > valQ[len(valQ)-1]{// keep biggest vals only
            indexQ = indexQ[:len(indexQ)-1]
            valQ = valQ[:len(valQ)-1]           
        }
        indexQ = append(indexQ,r)
        valQ = append(valQ,val)
        
        for r - indexQ[0] + 1 > k { 
            indexQ = indexQ[1:]
            valQ = valQ[1:]  
        }
        
        if r + 1 >= k{
            res = append(res,valQ[0])
        }
        
    }
    return res
}