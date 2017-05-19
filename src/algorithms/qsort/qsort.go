package qsort

func QuickSort(in []int, left, right int) []int {

	temp := in[left]
	p := left
	i, j := left, right
	for i <= j {
		if j >= p && in[j] >= temp {
			j--
		}
		if j >= p {
			in[p] = in[j]
			p = j
		}
		if in[i] <= temp&&i <= p {
			i++
		}
		if i<=p{
			in[p]=in[i]
			p=i
		}
	}
	in[p]=temp
	if p>left{
		QuickSort(in,left,p-1)
	}
	if(right>p){
		QuickSort(in,p+1,right)
	}
	return in
}
