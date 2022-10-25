func subsets(nums []int) [][]int {
    res :=  [][]int{}
    solution(nums,[]int{},&res)
    return res
}

func solution(nums,current []int, res *[][]int){
    fmt.Println(current,nums)
    *res = append(*res,current)
    for i,num := range nums{
        current = append(current,num)
        solution(nums[i+1:], current, res)
        current = append([]int{},current[:len(current)-1]...)
    }
}