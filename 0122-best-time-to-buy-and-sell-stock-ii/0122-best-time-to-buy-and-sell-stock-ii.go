func maxProfit(prices []int) int {
    profits := []int{}
   
    var l,r,current int
    
    for r <len(prices) {
        if prices[r] - prices[l] >= current {
            current = prices[r] - prices[l]
        }else{
            profits = append(profits,current)
            current = 0
            l = r
        }
        r++
    }
    var res int
    profits = append(profits,current)
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
