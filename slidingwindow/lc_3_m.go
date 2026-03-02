package slidingwindow

func lengthOfLongestSubstring(s string) int {
	window := make(map[byte]int)
	res := 0
	l, r := 0, 0
	for r < len(s) {
		curr := s[r]
		window[curr]++
		r++
		for window[curr] > 1 {
			window[s[l]]--
			l++
		}
		res = max(res, r-l)
	}
	return res
}
