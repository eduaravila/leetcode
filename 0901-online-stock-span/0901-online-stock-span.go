type StockSpanner struct {
    prices [][]int
}


func Constructor() StockSpanner {
    return StockSpanner{
        prices:[][]int{},
    }
}


func (this *StockSpanner) Next(price int) int {        
    res := 1
    i := len(this.prices)-1
    for i >= 0 {        
        if this.prices[i][0] > price{
            break
        }
        res += this.prices[i][1]
        i -= this.prices[i][1]
    }

    this.prices = append(this.prices,[]int{price,res})
    return res
}


/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */