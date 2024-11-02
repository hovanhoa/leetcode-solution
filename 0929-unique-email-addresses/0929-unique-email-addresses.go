func numUniqueEmails(emails []string) int {
	ans := make(map[string]bool)

	for i := 0; i < len(emails); i++ {
		var local []byte
        var j int
		for emails[i][j] != '@' && emails[i][j] != '+' {
			if emails[i][j] != '.' {
				local = append(local, emails[i][j])
			}
			j++
		}

		for emails[i][j] != '@' {
			j++
		}
		
		addr := string(local) + string(emails[i][j:])
		ans[addr] = true
	}
    
	return len(ans)
}