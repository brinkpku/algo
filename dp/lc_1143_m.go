package dp

// 自顶向下，递归+备忘录
func longestCommonSubsequence(text1 string, text2 string) int {
	var dfs func(i, j int) int
	mem := make([][]int, len(text1))
	for i := range mem {
		mem[i] = make([]int, len(text2))
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}
	dfs = func(i, j int) int {
		if i < 0 || j < 0 {
			return 0
		}
		if mem[i][j] != -1 {
			return mem[i][j]
		}
		if text1[i] == text2[j] {
			mem[i][j] = dfs(i-1, j-1) + 1
		} else {
			mem[i][j] = max(dfs(i-1, j), dfs(i, j-1))
		}
		return mem[i][j]
	}
	return dfs(len(text1)-1, len(text2)-1)
}

// 自底向上，dp table
func longestCommonSubsequence2(text1 string, text2 string) int {
	mem := make([][]int, len(text1)+1)
	for i := range mem {
		mem[i] = make([]int, len(text2)+1)
	}
	for i := 1; i <= len(text1); i++ {
		for j := 1; j <= len(text2); j++ {
			if text1[i-1] == text2[j-1] {
				mem[i][j] = mem[i-1][j-1] + 1
			} else {
				mem[i][j] = max(mem[i-1][j], mem[i][j-1])
			}
		}
	}
	return mem[len(text1)][len(text2)]
}

// 空间优化
func longestCommonSubsequence3(text1 string, text2 string) int {
	mem := make([]int, len(text2)+1)
	for i := 1; i <= len(text1); i++ {
		leftUp := mem[0]
		for j := 1; j <= len(text2); j++ {
			tmp := mem[j]
			if text1[i-1] == text2[j-1] {
				mem[j] = leftUp + 1
			} else {
				mem[j] = max(mem[j], mem[j-1])
			}
			leftUp = tmp
		}
	}
	return mem[len(text2)]
}
