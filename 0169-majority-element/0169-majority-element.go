func majorityElement(nums []int) int {
    inf := int(^uint(0)>>1)
    current,count := inf, 0
    for _, num := range nums{
        if num == current{
            count++
        }else if count == 0{
            count++
            current = num
        }else{
            count--
        }
    }
    return current
}