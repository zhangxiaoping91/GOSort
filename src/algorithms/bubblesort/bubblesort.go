package bubblesort

/**
冒泡排序
 */
func BubbleSort(in []int) [] int {
	for i, len := 0, len(in); i < len - 1; i++ {
		for j := 0; j < len - i - 1; j++ {
			if in[j] > in[j + 1] {
				in[j], in[j + 1] = in[j + 1], in[j]
			}
		}
	}
	return in;
}
