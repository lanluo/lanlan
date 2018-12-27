package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//fmt.Println(mergeSort(generateSlice(16)))
	mymap := make(map[string][]string, 0)
	myslice := []string{"hello", "world"}
	mymap["lan"] = myslice
	newslice := []string{"hi", "cathy"}
	fmt.Printf("address %p", &myslice[0])
	fmt.Println("")
	myslice = append(myslice, newslice...)
	//fmt.Println("")
	fmt.Printf("address %p", &myslice[0])
	fmt.Println("") //newnewslice := mymap["lan"]
	fmt.Printf("address mymap[] %p", &(mymap["lan"][0]))
	fmt.Println("")
	//mymap["lan"] = myslice
	fmt.Println(myslice)
	fmt.Println(mymap["lan"])
}

// func mergeSort(a []int) []int {
// 	if len(a) <= 1 {
// 		return a
// 	}
// 	m := len(a) / 2
// 	return merge(mergeSort(a[:m]), mergeSort(a[m:]))
// }

// func merge(a, b []int) []int {
// 	i := 0
// 	j := 0
// 	ret := []int{}
// 	for i < len(a) || j < len(b) {
// 		if i < len(a) && j < len(b) {
// 			if a[i] < b[j] {
// 				ret = append(ret, a[i])
// 				i++
// 			} else {
// 				ret = append(ret, b[j])
// 				j++
// 			}
// 		}
// 		if i >= len(a) && j < len(b) {
// 			tmp := b[j:]
// 			ret = append(ret, tmp...)
// 			break
// 		}
// 		if j >= len(b) && i < len(a) {
// 			tmp := a[i:]
// 			ret = append(ret, tmp...)
// 			break
// 		}
// 	}
// 	return ret
// }

// func mergeSort(a []int, l, h int) []int {
// 	if l == h {
// 		return []int{a[l]}
// 	}
// 	if l+1 == h {
// 		if a[l] > a[h] {
// 			return []int{a[h], a[l]}
// 		}
// 		return []int{a[l], a[h]}
// 	}

// 	m := (l + h) / 2
// 	leftSlice := a[l : m+1]
// 	rightSlice := a[m+1 : h+1]
// 	sortedLeftSlice := mergeSort(leftSlice, l, m)
// 	sortedRightSlice := mergeSort(rightSlice, m+1, h)
// 	return mergeSortList(sortedLeftSlice, sortedRightSlice)
// }
// func mergeSortList(a, b []int) []int {
// 	i := 0
// 	j := 0
// 	slice := make([]int, 0)
// 	for i < len(a) && j < len(b) {
// 		if i < len(a) && j < len(b) {
// 			if a[i] < b[j] {
// 				slice = append(slice, a[i])
// 				i++
// 			} else {
// 				slice = append(slice, b[j])
// 				j++
// 			}
// 		}
// 		if i >= len(a) && j < len(b) {
// 			tmp := b[j:]
// 			slice = append(slice, tmp...)
// 			break
// 		}
// 		if j >= len(b) && i < len(a) {
// 			tmp := a[i:]
// 			slice = append(slice, tmp...)
// 			break
// 		}
// 	}
// 	fmt.Println(slice)
// 	return slice
// }

func generateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(2*i + 32)
	}

	//fmt.Println(slice)
	return slice
}
