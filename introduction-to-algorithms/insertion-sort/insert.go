package isort

func isort(s []int) []int {
	for nCur := 1; nCur < len(s); nCur++ {
		key := s[nCur]
		nPrev := nCur - 1
		for nPrev >= 0 && s[nPrev] > key {
			s[nPrev+1] = s[nPrev]
			nPrev--
		}
		s[nPrev+1] = key
	}
	return s
}

func isortSlices(s [][]int) [][]int {
	sortedResults := make([][]int, len(s))
	for _, slizzle := range s {
		sortedResults = append(sortedResults, isort(slizzle))
	}
	return sortedResults
}

func isortSlicesConcurrently(s [][]int) [][]int {
	sortedResults := make([][]int, len(s))
	results := make(chan []int)

	for _, slizzle := range s {
		go func(s []int) {
			results <- isort(s)
		}(slizzle)
	}

	for i := 0; i < len(sortedResults); i++ {
		sorted := <-results
		sortedResults[i] = sorted
	}
	return sortedResults
}

// func isortSlicesConcurrent(s [][]int) [][]int {
// 	sLen := len(s)
// 	sMid := int(math.Floor(float64(sLen) / 2))
// 	s1 := s[:sMid]
// 	s2 := s[sMid:]
// 	sortedResults

// 	for _, slizzle := range s {
// 		go func() {
// 			sl = isort(s)
// 		}()
// 	}

// 	for nCur := 1; nCur < len(s); nCur++ {
// 		key := s[nCur]
// 		nPrev := nCur - 1
// 		for nPrev >= 0 && s[nPrev] > key {
// 			s[nPrev+1] = s[nPrev]
// 			nPrev--
// 		}
// 		s[nPrev+1] = key
// 	}
// 	return s
// }
