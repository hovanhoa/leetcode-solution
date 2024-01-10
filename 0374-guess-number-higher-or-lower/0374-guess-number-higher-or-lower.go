/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
    left := 1
    right := n

    for {
        num := (left + right)/2
        if guess(num) == 0 {
            return num
        } else if guess(num) == 1 {
            left = num + 1
        } else {
            right = num - 1
        }
    }
}