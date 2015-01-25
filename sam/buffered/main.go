package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	inputs := make(chan int, 100)

	go func() {
		for {
			var input string
			fmt.Scanln(&input)
			num, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("err:", err)
			} else {
				inputs <- num
			}
		}
	}()

	timer := time.NewTimer(time.Second * 5)

	sum := 0

	for {
		select {
		case num := <-inputs:
			sum += num
			timer.Reset(time.Second * 5)
		case <-timer.C:
			fmt.Println(sum)
			return
		}
	}
}
