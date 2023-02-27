type NumArray struct {
    pref,nums []int
}


func Constructor(nums []int) NumArray {
    pref := []int{}
    var sum int
    for _, num := range nums{
        sum += num
        pref = append(pref,sum)
    }
    return NumArray{
        nums: nums,
        pref: pref,
    }
}


func (this *NumArray) SumRange(left int, right int) int {
    return this.pref[right] - this.pref[left] + this.nums[left]
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */