type NumArray struct {
   bitree []int 
}

func Constructor(nums []int) NumArray {
    numArray := NumArray {make([]int, len(nums)+1)}
    
    for i, v := range nums {
        numArray.Update(i, v)
    }
    
    return numArray
}

func (this *NumArray) Update(index int, val int)  {
    diff := val-this.SumRange(index, index)
    index++
    
    for index < len(this.bitree) {
        this.bitree[index] += diff
        index += index&-index
    }
}

func (this *NumArray) SumRange(left int, right int) int {  
    return this.Sum(right)-this.Sum(left-1)
}

func (this *NumArray) Sum(index int) int {
    index++
    result := 0
    
    for index > 0 {
        result += this.bitree[index]
        index -= index&-index
    }
    
    return result
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := ConstructorNumArray(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */

