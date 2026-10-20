func mySqrt(x int) int {
    p, q := 1, x

    for p <= q {
        mid := (p+q)/2
        if mid * mid == x {
            return mid
        }
        if mid * mid < x {
            p = mid + 1
        } else {
            q = mid - 1
        }
    }

    return q

}
