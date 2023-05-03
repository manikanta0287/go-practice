package logics

func Unique(arr []int) []int {
	occured := map[int]bool{}
	result := []int{}

	for e := range arr {
		if occured[arr[e]] != true {
			occured[arr[e]] = true

			result = append(result, arr[e])
		}
	}
	return result
}
