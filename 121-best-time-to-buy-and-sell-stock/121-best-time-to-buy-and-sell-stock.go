func maxProfit(prices []int) int {
    l,res := 0,0
    for i:=1 ;i<len(prices); i++{
        profit := prices[i] - prices[l]
        if profit > res {
            res= profit
        }
        if prices[i] < prices[l] {
            l = i
        }        
    }
    
    return res
}