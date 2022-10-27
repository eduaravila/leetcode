func majorityElement(nums []int) int {
    sort.Ints(nums)
    var current, count int
    avg := len(nums)/ 2
    for _, num := range nums{
        if current != num {
            current = num
            count=0
        }
        count++
        if count > avg{
            return num
        }
    }
    return 0
}