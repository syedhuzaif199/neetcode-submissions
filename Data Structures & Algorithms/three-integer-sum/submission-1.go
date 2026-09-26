import "slices"
func threeSum(nums []int) [][]int {
    slices.Sort(nums)
    triplets := [][]int{}
    i := 0
    for i < len(nums) - 1 {
        p, q := i + 1, len(nums)-1
        for p < q {
            val := nums[i] + nums[p] + nums[q]
            skipP, skipQ := true, true
            if val == 0 {
                triplets = append(triplets, []int{nums[i], nums[p], nums[q]})
            } else if val < 0 {
                skipQ = false
            } else {
                skipP = false
            }
            if skipP {
                x := nums[p]
                for p < q && nums[p] == x {
                    p += 1
                }
            }
            if skipQ {
                x := nums[q]
                for q > p && nums[q] == x {
                    q -= 1
                }
            }
        }

        x := nums[i]
        i+=1
        for i < len(nums) && nums[i] == x {
            i += 1
        }
    }
    return triplets
}
