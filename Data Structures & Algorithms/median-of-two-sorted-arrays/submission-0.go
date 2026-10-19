func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n := len(nums1) + len(nums2)

	skips := n/2 + n%2 - 1

	p, q := 0, 0

	for _ = range skips {
		if p < len(nums1) && q < len(nums2) {
			if nums1[p] < nums2[q] {
				p += 1
			} else {
				q += 1
			}
		} else if p < len(nums1) {
			p += 1
		} else if q < len(nums2) {
			q += 1
		} else {
			// unreachable for valid test cases
		}
	}

	median := 0

	if p < len(nums1) && q < len(nums2) {
		if nums1[p] < nums2[q] {
			median = nums1[p]
			if n % 2 == 0 {
				if p + 1 < len(nums1) && nums1[p+1] < nums2[q] {
					median += nums1[p+1]
				} else {
					median += nums2[q]
				}
			}
		} else {
			median = nums2[q]
			if n % 2 == 0 {
				if q + 1 < len(nums2) && nums2[q+1] < nums1[p] {
					median += nums2[q+1]
				} else {
					median += nums1[p]
				}
			}
		}
	} else if p < len(nums1) {
		median = nums1[p]
		if n % 2 == 0 {
			median += nums1[p+1]
		}
	} else if q < len(nums2) {
		median = nums2[q]
		if n % 2 == 0 {
			median += nums2[q+1]
		}
	}

	if n % 2 == 0 {
		return float64(median)/2
	}
	return float64(median)

}
