func largestNumber(nums []int) string {
    res := []string{}
    sort.Slice(nums, func(a,b int)bool{
        sa := fmt.Sprintf("%d%d",nums[a],nums[b])
        sb := fmt.Sprintf("%d%d",nums[b],nums[a])
        ia,_:= strconv.Atoi(sa)
        ib,_ := strconv.Atoi(sb)
        return ia > ib 
    })
    
    for _,num := range nums{
        res = append(res,fmt.Sprintf("%d",num))
    }
    
    result := strings.TrimLeft(strings.Join(res,""),"0")
    if result < "0"{
        return "0"
    }
    
    return result
}