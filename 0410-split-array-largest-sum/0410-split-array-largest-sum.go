func getMax(vals ...int)int{
    var res int
    for _,val := range vals{
        if val > res{
            res = val
        }
    }
    return res
}

func splitArray(nums []int, k int) int {
    
    var sum,res int
    
    for _, num := range nums{
        sum+=num
    }
    
    l,r := getMax(nums...),sum
    
    for l <= r {
        m := l + ((r-l) /2)
        
        if canSplit(nums, m, k){
            res = m
            r = m-1
        }else{
            l = m+1
        }
    }
    
    return res
}

func canSplit(nums []int, sum, k int)bool{
    groups := []int{}
    
    for _,num := range nums{
        if len(groups)<1 || (groups[len(groups)-1] +num) > sum {
            groups = append(groups,0)
        }
        groups[len(groups)-1]+=num
        
    }
    
    return len(groups) <=k
}