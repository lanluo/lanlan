package main

func insertsort(a []int) []int {
	if len(a) <= 1 {
		return a
	}
	for i := 1; i < len(a); i++ {
		j := i
		for j > 0 && a[j] < a[j-1] {
			a[j], a[j-1] = a[j-1], a[j]
			j--
		}
	}
	//fmt.Println(a)
	return a
}
