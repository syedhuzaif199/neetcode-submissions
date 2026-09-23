import "slices"
func longestConsecutive(nums []int) int {
    slices.Sort(nums)
    longest := 0
    count := 0
    var last int
    i := 0
    for i < len(nums) {
        num := nums[i]
        if count == 0 {
            count = 1
        } else if num == last + 1 {
            count += 1
        } else if num != last {
            count = 0
            i -= 1
        }
        last = num
        longest = max(longest, count)
        i += 1
    }

    return longest
}