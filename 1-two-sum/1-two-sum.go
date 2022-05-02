func twoSum(nums []int, target int) []int {
    // add the numbers to a hash table
    // iterate through the array and substract current - total
    // check if the remaining exist in the hashtable
    
    counter := make(map[int]int)
    for idx,v := range nums{ // O(n)
        remaining := target - v 
        if _,e := counter[remaining]; e {
            return []int{idx,counter[remaining]}
        }
        counter[v]= idx 
        
    }
    

    return []int{}
}