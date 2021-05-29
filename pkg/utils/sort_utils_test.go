package utils

import (
	"testing"
)


func TestBubbleSort(t *testing.T) {
	t.Run("given unsorted arr, return sorted arr with same number of elements", func(t *testing.T) {
		arr := []int{2, 5, 1, 7, 4, 9, 2, 3, 5, 1, 10}

		got := BubbleSort(arr)
	
		want := []int{1, 1, 2, 2, 3, 4, 5, 5, 7, 9, 10}

		if len(got) != len(arr) {
			t.Fatalf("did not get back same number of arr elements")
		}
	
		AssertEqualInstance(t, got, want)
	})

	t.Run("given nil input arr, return []", func(t *testing.T) {
		got := BubbleSort(nil)
		want := []int{}

		AssertEqualInstance(t, got, want)
	})
}

func BenchmarkBubbleSort(b *testing.B) {
	arr := GenerateIntSlice(1000)

	for i := 0; i < b.N; i++ {
		BubbleSort(arr)
	}
}

func TestMergeSort(t *testing.T) {
	t.Run("sorts and returns arr with same length", func(t *testing.T) {
		arr := []int{2, 5, 1, 7, 4, 9, 2, 3, 5, 1, 10}

		got := MergeSort(arr)
	
		want := []int{1, 1, 2, 2, 3, 4, 5, 5, 7, 9, 10}
	
		if len(got) != len(arr) {
			t.Fatalf("did not get back same number of arr elements")
		}
	
		AssertEqualInstance(t, got, want)
	})

	t.Run("returns [] for nil input", func(t *testing.T) {
		got := MergeSort(nil)
		want := []int{}

		AssertEqualInstance(t, got, want)
	})
}

func BenchmarkMergeSort(b *testing.B) {
	arr := GenerateIntSlice(1000)

	for i := 0; i < b.N; i++ {
		MergeSort(arr)
	}
}