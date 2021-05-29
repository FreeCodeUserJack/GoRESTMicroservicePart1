package main

import (
	"fmt"

	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/app"
)

func main() {
	fmt.Println("entry point")
	app.StartApp()

	// arr := []int{2, 5, 1, 7, 4, 9, 2, 3, 5, 1, 10}
	// fmt.Println(utils.BubbleSort(arr))
	// fmt.Println(utils.MergeSort(arr))
}