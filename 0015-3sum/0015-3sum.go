func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    res := [][]int{}
    for i,num := range nums{
        if i > 0 && nums[i-1] == nums[i]{
            continue
        }
        l,r := i+1, len(nums)-1
        for l < r{
            total := num + nums[l] + nums[r]
            if total == 0{
                res = append(res,[]int{num,nums[l],nums[r]})
                r--
                for l < r && nums[r] == nums[r+1]{
                    r--
                }
            }else if total > 0 {
                r--
            }else{
                l++
            }
        }
    }
    return res
}