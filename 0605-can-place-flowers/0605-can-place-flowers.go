func canPlaceFlowers(flowerbed []int, n int) bool {
    for idx, val := range flowerbed {
        if val == 0 && (idx == 0 || flowerbed[idx-1] == 0) && (idx == len(flowerbed) - 1 || flowerbed[idx+1] == 0) {
            n -= 1
            flowerbed[idx] = 1
        }
    }

    return n <= 0
}