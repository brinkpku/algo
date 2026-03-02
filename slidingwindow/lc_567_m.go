package slidingwindow

func checkInclusion(s1 string, s2 string) bool {
	target := [26]int{}
	for _, r := range s1 {
		target[r-'a']++
	}
	window := [26]int{}
	l, r := 0, 0
	for r < len(s2) {
		charIdx := s2[r] - 'a'
		r++
		window[charIdx]++
		for l < r && window[charIdx] > target[charIdx] {
			window[s2[l]-'a']--
			l++
		}
		if window == target {
			return true
		}
	}
	return false
}
