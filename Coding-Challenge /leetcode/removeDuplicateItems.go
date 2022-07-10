func removeDuplicates(nums []int) int {
    if l := len(nums); l <= 1 {
        return l
    }
 
    i := 1
    for _, x := range nums[1:] {
        if x != nums[i-1] {
            nums[i] = x
            i++
        }
    }
    return i
}

