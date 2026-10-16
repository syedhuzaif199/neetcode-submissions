import "slices"
func numRescueBoats(people []int, limit int) int {
    slices.Sort(people)
    p, q := 0, len(people)-1
    count := 0
    for p <= q {
        count += 1
        if people[p] + people[q] <= limit {
            p += 1
        }
        q -= 1
    }

    return count
}
