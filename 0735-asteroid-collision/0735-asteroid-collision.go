func asteroidCollision(asteroids []int) []int {
    ans := make([]int, 0)
    for i := 0; i < len(asteroids); i++ {
        if asteroids[i] > 0 {
            ans = append(ans, asteroids[i])
        } else {
            if len(ans) == 0 {
                ans = append(ans, asteroids[i])
                continue
            }

            if ans[len(ans)-1] == -asteroids[i] {
                ans = ans[:len(ans)-1]
            } else if ans[len(ans)-1] < -asteroids[i] {
                ans = ans[:len(ans)-1]
                i--
            }
        }
    }

    return ans
}