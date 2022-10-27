func majorityElement(nums []int) int {
    current,count := nums[0],1
    for _, num := range nums[1:]{
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