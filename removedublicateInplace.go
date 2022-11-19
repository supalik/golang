
func removeDuplicates(nums []int) int {
    newlen:=1
    for i:=1; i< len(nums); i++{
        if nums[i-1] != nums[i]{
            nums[newlen]=nums[i]
            newlen++
        }
    }
    return newlen
}
