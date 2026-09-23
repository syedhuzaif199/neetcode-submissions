type Range struct {
	start int
	end int
}

func longestConsecutive(nums []int) int {
	m := map[int]*Range{}
	longest := 0
	for _, num := range nums {
		r, _ := m[num]
		if r != nil {
			continue
		}

		pre, _ := m[num-1]
		post, _ := m[num+1]
		if pre != nil && post != nil && num == pre.end + 1 && num == post.start - 1 {
			m[pre.start] = nil
			m[pre.end] = nil
			m[post.start] = nil
			m[post.end] = nil
			r := Range{
				start: pre.start,
				end: post.end,
			}
			m[r.start] = &r
			m[r.end] = &r
			longest = max(longest, r.end - r.start + 1)
		} else if pre != nil && num == pre.end + 1 {
			m[pre.end] = nil
			pre.end = num
			m[pre.start] = pre
			m[pre.end] = pre
			longest = max(longest, pre.end - pre.start + 1)
		} else if post != nil && num == post.start - 1 {
			m[post.start] = nil
			post.start = num
			m[post.start] = post
			m[post.end] = post
			longest = max(longest, post.end - post.start + 1)
		} else if post == nil && pre == nil {
			r := Range{
				start: num,
				end: num,
			}
			m[r.start] = &r
			m[r.end] = &r
			longest = max(longest, r.end - r.start + 1)

		}
	}
	return longest
}
