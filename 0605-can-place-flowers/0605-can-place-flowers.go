func canPlaceFlowers(flowerbed []int, n int) bool {
    num := 0 
    for i := 0; i < len(flowerbed); i++ {
        if flowerbed[i] == 0 && (i == 0 || flowerbed[i-1] == 0) && (i == len(flowerbed) - 1 || flowerbed[i+1] == 0) {
            num += 1
            flowerbed[i] = 1
        }
    }

    return num >= n
    
}