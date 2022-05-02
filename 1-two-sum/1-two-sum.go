func twoSum(nums []int, target int) []int {
    // add the numbers to a hash table
    // iterate through the array and substract current - total
    // check if the remaining exist in the hashtable
    
    counter := make(map[int]int)
    for idx,v := range nums{ // O(n)
        counter[v]= idx
    }
    
    for idx,v := range nums{ //O(n)
        remaining := target - v 
        if ridx,e := counter[remaining]; e && ridx != idx {
            return []int{idx,counter[remaining]}
        }
    }
    return []int{}
}