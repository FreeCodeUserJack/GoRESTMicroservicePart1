package main

import "fmt"

func main() {
	done := make(chan int)
	go helloWorld(done)
	<- done

	c := make(chan string)

	go func(input chan string) {
		c <- "hi"
	}(c)

	res := <-c
	fmt.Println(res)

	cr := make(chan int, 5)

	go func(input chan int) {
		for i := 0; i < 5; i++ {
			cr <- i
		}
		close(cr)
	}(cr)

	for num := range cr {
		fmt.Println(num)
	}
}

func helloWorld(done chan int) {
	fmt.Println("Hello world")
	done <- 1
}