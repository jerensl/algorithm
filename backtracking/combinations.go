package backtracking

func Combinations(n, k int) [][]int {
	var allComb [][]int
	
	comb := make([]int, 0, k)

	dfs(1, n, k, comb, &allComb)
	
	return allComb
}

func dfs(start, n, k int, comb []int, allComb *[][]int) {
	if k == 0 {
		cp := make([]int, len(comb))

		copy(cp, comb)

		*allComb = append(*allComb, cp)
		return
	}

	for i := start; i <= n-k+1; i++ {
		dfs(i+1, n, k-1, append(comb, i), allComb)
	}
}