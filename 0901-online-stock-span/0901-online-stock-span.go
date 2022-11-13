type StockSpanner struct {
    prices []int    
}


func Constructor() StockSpanner {
    return StockSpanner{
        prices:[]int{},
    }
}


func (this *StockSpanner) Next(price int) int {
    res := 1
    for i:= len(this.prices)-1 ; i >= 0; i--{
        if this.prices[i] > price{
            break
        }
        res++
    }    
    this.prices = append(this.prices,price) 
    return res
}


/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */