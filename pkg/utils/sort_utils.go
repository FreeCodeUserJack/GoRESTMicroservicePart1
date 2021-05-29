package utils

func BubbleSort(arr []int) []int {
	if arr == nil {
		return []int{}
	}

	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			if arr[j] < arr[i] {
				buf := arr[j]
				arr[j] = arr[i]
				arr[i] = buf
			}
		}
	}
	return arr
}

func MergeSort(arr []int) []int {
	if arr == nil {
		return []int{}
	}

	var num = len(arr)

	if num == 1 {
		return arr
	}

	middle := int(num / 2)

	var (
		left  = make([]int, middle)
		right = make([]int, num-middle)
	)

	for i := 0; i < num; i++ {
		if i < middle {
			left[i] = arr[i]
		} else {
			right[i-middle] = arr[i]
		}
	}
	return merge(MergeSort(left), MergeSort(right))
}

func merge(left, right []int) (result []int) {
	result = make([]int, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}

	for k := 0; k < len(right); k++ {
		result[i] = right[k]
		i++
	}

	return
}