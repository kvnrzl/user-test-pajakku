package main

import "sort"

func removeDuplicatesAndSorting(values []int) []int {
	// menghapus data yang duplikat
	uniqueValues := make(map[int]bool)
	for _, value := range values {
		uniqueValues[value] = true
	}

	// menambahkan data yang unik ke dalam slice results
	var results []int
	for key := range uniqueValues {
		results = append(results, key)
	}

	// melakukan sorting secara desecending
	sort.Slice(results, func(i, j int) bool {
		return results[i] > results[j]
	})

	return results
}
