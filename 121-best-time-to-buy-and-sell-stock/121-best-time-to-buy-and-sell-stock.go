func getMax(a,b int)int{
    if a > b {
        return a
    }
    return b
}

func maxProfit(prices []int) int {
    if len(prices) < 2 {
        return 0 
    }
    
    buy, sell := 0,1
    var res int
    for sell < len(prices){
        
        if prices[sell] < prices[buy]{
            buy = sell
        }else{
            res = getMax(res, prices[sell] - prices[buy])
        }
        sell++
    }
    
    
    
    return res
}