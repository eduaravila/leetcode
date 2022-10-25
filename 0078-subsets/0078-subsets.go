func subsets(nums []int) [][]int {
    res :=  [][]int{}
    solution(nums,[]int{},0,&res)
    return res
}

func solution(nums,current []int,index int, res *[][]int){
    
    *res = append(*res,current)
    for index < len(nums) {
        current = append(current,nums[index])
        solution(nums, current,index+1, res)
        current = append([]int{},current[:len(current)-1]...)
        index++
    }
}