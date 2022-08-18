package main

import "fmt"

func b() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
			ch2 <- i
		}
	}()
	for i := 0; i < 15; i++ {
		select {
		case x := <-ch1:
			fmt.Printf("receive %d from channel 1-1\n", x)
		case x := <-ch1:
			fmt.Printf("receive %d from channel 1-2\n", x)
		case y := <-ch2:
			fmt.Printf("receive %d from channel 2\n", y)
		}
	}
}

func main() {
	b()
}
