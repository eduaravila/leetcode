func sortColors(nums []int)  { 
    sortq(nums,0,len(nums)-1)
}


func partition(nums []int, low,high int)int{
    pi := nums[high]
    i := low-1
    for low < high{
        if nums[low] < pi {
            i++
            nums[i],nums[low] = nums[low],nums[i]
        }
        low++
    }
    nums[i+1] ,nums[high] = nums[high] ,nums[i+1]
    return i +1
       
}

func sortq(nums []int, l,r int){
    if l > r{
        return 
    }
    
    pi :=partition(nums,l,r)
    sortq(nums,l,pi-1)
    sortq(nums,pi+1,r)
}