func canPlaceFlowers(flowerbed []int, n int) bool {
    for i := 0; i < len(flowerbed); i++ {
        if flowerbed[i] == 0 && (i == 0 || flowerbed[i-1] == 0) && (i == len(flowerbed) - 1 || flowerbed[i+1] == 0) {
            flowerbed[i] = 1
            n -= 1
        }
    }

    return n <= 0
}