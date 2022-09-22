func solutionSplit(nums []int, l, r int) int {
    pivot :=  r
    i := l
    for _, num := range nums[l:r]{
      if num <= nums[pivot]{
        nums[l],nums[i] = nums[i], nums[l]
        l++
        }
        i++
    }
    nums[l], nums[pivot] = nums[pivot], nums[l]
    return l
}

func findKthLargest(nums []int, k int) int {
    
    l, r := 0,len(nums)-1
    k = len(nums) - k
    for l < r {
        pivot := solutionSplit(nums,l , r)
        
        if k < pivot {
            r = pivot - 1
        }else if k > pivot{
            l = pivot + 1
        }else{
            break // if k == pivot we are sure that the nums[pivot] value is the k largerst
        }
    }
    return nums[k]
}