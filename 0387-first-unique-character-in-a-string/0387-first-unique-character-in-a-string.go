func firstUniqChar(s string) int {
	// This is a constant space allocation (ie: always 26)
	var counts = make([]int, 26)
	
	// Insert all the character's count into counts array
	for i := 0; i < len(s); i++ {
		counts[s[i] - 'a']++
	}

	// Find the first element with count 1 
    for i := 0; i < len(s); i++ {
		if counts[s[i] - 'a'] == 1 {
			return i
		}
	}

	return -1
}