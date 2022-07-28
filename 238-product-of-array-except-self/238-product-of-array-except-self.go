func productExceptSelf(nums []int) []int {
    lr := []int{}
    rl := []int{}
    lr = append(lr,nums...)
    rl = append(rl,nums...)
    l,r := 0,1
    for r < len(nums)-1{
        lr[r] = lr[l] * lr[r]
        l++
        r++
    }
    
    l,r = len(nums)-1,len(nums)-2
    for r > 0{
        rl[r] = rl[l] * rl[r]
        l--
        r--
    }
    
    res := make([]int,len(nums))
    fmt.Println(lr,rl)
    l = 0
    for l < len(nums){
        if l-1 < 0 {
            res[l] = rl[l+1]
        } else if l+1 >= len(nums){
            res[l] = lr[l-1]
        }else{
            res[l] = lr[l-1] * rl[l+1]        
        
        }
        l++
    }
   
    return res 
}