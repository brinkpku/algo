package slidingwindow

import "math"

func countCompleteSubstrings(word string, k int) int {
	ans := 0
	for i := range 26 {
		valid := 0
		window := make(map[byte]int)
		for left, right := 0, 0; right < len(word); {
			window[word[right]]++
			if window[word[right]] == k {
				valid++
			}
			right++
			// pay attention to uint8(byte) overflow when result is negative
			if right > 1 && int(math.Abs(float64(word[right-1])-
				float64(word[right-2]))) > 2 {
				for left < right-1 {
					if window[word[left]] == k {
						valid--
					}
					window[word[left]]--
					left++
				}
			}
			if right-left > k*(i+1) {
				if window[word[left]] == k {
					valid--
				}
				window[word[left]]--
				left++
			}
			if valid == i+1 {
				ans++
			}
		}
	}
	return ans
}
