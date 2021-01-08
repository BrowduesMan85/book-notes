package isort

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestInsertionSort(t *testing.T) {
	s := []int{3, 2, 1}

	want := []int{1, 2, 3}
	got := isort(s)

	assert.Equal(t, want, got)
}

func BenchmarkInsertionSort(b *testing.B) {
	rand.Seed(time.Now().Unix())

	for i := 0; i < b.N; i++ {
		testData := rand.Perm(1000)
		isort(testData)
	}
}

func BenchmarkInsertionSortSlices(b *testing.B) {
	rand.Seed(time.Now().Unix())

	testData := make([][]int, 100)
	for i := 0; i < len(testData); i++ {
		testData[i] = rand.Perm(1000)
	}

	for i := 0; i < b.N; i++ {
		isortSlices(testData)
	}
}

func BenchmarkInsertionSortSlicesConcurrently(b *testing.B) {
	rand.Seed(time.Now().Unix())

	testData := make([][]int, 100)
	for i := 0; i < len(testData); i++ {
		testData[i] = rand.Perm(1000)
	}

	for i := 0; i < b.N; i++ {
		isortSlicesConcurrently(testData)
	}
}
