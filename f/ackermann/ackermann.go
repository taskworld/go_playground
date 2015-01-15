package main

import "fmt"

func main() {
	// fmt.Println(ackermann(4, 1))

	msg := make(chan int)

	go func() {
		msg <- ackermann(4, 1)
	}()

	fmt.Println(<-msg)
}

func ackermann(m int, n int) (result int) {
	if m == 0 {
		result = n + 1
	} else if n == 0 {
		result = ackermann(m-1, 1)
	} else {
		result = ackermann(m-1, ackermann(m, n-1))
	}
	return
}
