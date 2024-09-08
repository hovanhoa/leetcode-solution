func abs(a int) int {
    if a < 0 {
        return -a
    }

    return a
}

func asteroidCollision(asteroids []int) []int {
    var stack []int
    for i := 0; i < len(asteroids); i++ {
        next := asteroids[i]
        if len(stack) == 0 {
            stack = append(stack, next)
            continue
        }

        last := stack[len(stack)-1]
        if last < 0 || next > 0 {
            stack = append(stack, next)
            continue
        }

        absNext := abs(next)
        if (last > absNext) {
            continue
        } else if (absNext > last) {
            stack = stack[:len(stack)-1]
            i--
        } else {
            stack = stack[:len(stack)-1]
        }
    }

    return stack
}