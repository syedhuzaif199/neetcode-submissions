import "slices"

type pair struct {
	x, y int
}
func carFleet(target int, position []int, speed []int) int {
	groups := make([]pair, len(position))
	for i := range position {
		groups[i] = pair{position[i], speed[i]}
	}

	slices.SortFunc(groups, func(a, b pair) int {
		return b.x - a.x
	})

	count := 0
	prev := float64(0)

	for _, group := range groups {
		dx := target - group.x
		dt := float64(dx) / float64(group.y)

		if dt > prev {
			prev = dt
			count += 1
		}
	}

	return count
}
