func canPlaceFlowers(flowerbed []int, n int) bool {
	for i, plot := range flowerbed {
		if plot == 1 {
			continue
		}

		if (i == 0 || flowerbed[i-1] == 0) && (i == len(flowerbed)-1 || flowerbed[i+1] == 0) {
			flowerbed[i] = 1
			n--
		}
	}

	return n <= 0
}