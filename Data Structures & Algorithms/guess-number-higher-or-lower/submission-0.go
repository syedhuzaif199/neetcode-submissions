/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
    p, q := 1, n

    for p <= q {
        mid := (p+q)/2
        ans := guess(mid)
        if ans == 0 {
            return mid
        }
        if ans < 0 {
            q = mid - 1
        } else {
            p = mid + 1
        }
    }

    return 0
 
}
