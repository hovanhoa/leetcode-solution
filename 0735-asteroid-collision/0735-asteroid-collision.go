func asteroidCollision(asteroids []int) []int {
    stack := []int{}
    for i := 0; i < len(asteroids); i++ {
        if asteroids[i] > 0 {
            stack = append(stack, asteroids[i])
        } else {
            if len(stack) == 0 || stack[len(stack)-1] < 0 {
                stack = append(stack, asteroids[i])
                continue
            }

            if stack[len(stack)-1] == -asteroids[i] {
                stack = stack[:len(stack)-1]
            } else if stack[len(stack)-1] < -asteroids[i] {
                stack = stack[:len(stack)-1]
                i -= 1
            }
        }
    }

    return stack
}