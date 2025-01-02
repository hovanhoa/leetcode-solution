func isRobotBounded(instructions string) bool {
    dirX, dirY := 0, 1
    x, y := 0, 0

    for _, v := range instructions {
        if v == 'G' {
            x, y = x + dirX, y + dirY
        } else if v == 'L' {
            dirX, dirY = -1 * dirY, dirX
        } else {
            dirX, dirY = dirY, -1 * dirX
        }
    }

    return (x == 0 && y == 0) || dirX != 0 || dirY != 1
}