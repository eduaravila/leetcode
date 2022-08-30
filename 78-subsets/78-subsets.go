func subsets(nums []int) [][]int {
    res := [][]int{}
    solution(nums,[]int{},&res,0)
    return append(res,[]int{})
}

func solution(nums, subset []int,res *[][]int,current int) {
    if current >= len(nums){
        return 
    }

    for i := current ; i < len(nums) ; i++ {
        num := nums[i]
        subset = append(subset,num)
        *res = append(*res, subset)
        solution(nums,subset,res,i+1)                
        subset = append([]int{},subset[:len(subset)-1]...)
    }    
    
}