import "slices"
func topKFrequent(nums []int, k int) []int {
    m := map[int]int{}
    for _, num := range nums {
        m[num] += 1
    }
    keys := make([]int, 0, len(m))
    for key := range m {
        keys = append(keys, key)
    }

    slices.SortFunc(keys, func(a, b int) int {
        return m[b] - m[a]
    })

    return keys[:k]
}