func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    res := [][]int{}
    solution(nums,[]int{},0,3,0,&res)
    return res
}

func solution(nums,current []int, index,k,total int, res *[][]int){
    if k == 2{
        l,r := index, len(nums)-1
        for l < r{
            sum := total + nums[l] + nums[r]
            if sum == 0{
                *res = append(*res, append([]int{nums[l],nums[r]},current...))
                r--
                for l < r && nums[r] == nums[r+1]{
                    r--
                }
            }else if sum > 0 {
                r--
            }else{
                l++
            }
        }
        return 
    }
    for i := index ; i < len(nums); i++{
        if i > index && nums[i-1] == nums[i]{
            continue
        }
        current = append(current,nums[i])
        solution(nums,current,i+1,k-1,total+ nums[i],res)
        current = append([]int{},current[:len(current)-1]...)
    }
}