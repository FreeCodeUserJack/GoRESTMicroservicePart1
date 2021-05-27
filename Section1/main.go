package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("hello world")

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func Add(a, b int) int {
	return a + b
}