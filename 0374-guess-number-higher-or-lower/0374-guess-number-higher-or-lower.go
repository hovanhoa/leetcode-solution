/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
    l, r := 0, n
    for l <= r {
        pick := (l+r)/2
        if guess(pick) == 0 {
            return pick
        } else if guess(pick) == 1 {
            l = pick + 1
        } else {
            r = pick - 1
        }
    }

    return -1
}