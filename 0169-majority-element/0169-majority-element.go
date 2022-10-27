func majorityElement(nums []int) int {
    avg := len(nums) / 2
    count := make(map[int]int)
    for _, num := range nums{
        count[num]++
        if count[num] > avg{
            return num
        }
    }
    return 0
}