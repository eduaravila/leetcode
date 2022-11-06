func moveZeroes(nums []int)  {
    var zero int
    
    for _, num := range nums {
        if num != 0{
          nums[zero] = num
          zero++
        } 
    }
    
    for zero < len(nums){
        nums[zero] = 0
        zero++
    }
}