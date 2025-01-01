package leetcode

func maximumPopulation(logs [][]int) int {
	arr := make([]int, 101)
	for i := 0; i < len(logs); i++ {
		start := logs[i][0]
		end := logs[i][1] - 1
		for j := start; j <= end; j++ {
			arr[j-1950]++
		}
	}
	max := arr[0]
	y := 0
	for j := 0; j < len(arr); j++ {
		if arr[j] > max {
			y = j
			max = arr[j]
		}
	}
	return y + 1950
}
