func maxProfit(prices []int) int {
    profits := make([]int,len(prices))
   
    var l,r int
    
    for r <len(prices) {
        if prices[r] - prices[l] >= profits[l] {
            profits[l] =prices[r] - prices[l]
        }else{
            l = r
         
        }
        r++
    }
    var res int
    for _,profit := range profits{
        res+=profit
    }
    
    return res
}

// 1,2,3,4,5
// 0,0,0,0,0
// 0,1,2,3,4

// 7,1,5,3,6,4

// min 7,1,1,1
// max 7,1,5,3

// 0,

// if current < max
// update both min and max
