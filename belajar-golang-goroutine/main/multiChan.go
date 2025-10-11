package main

import "fmt"

func getAverage(numbers []int, ch chan float64) {
	sum := 0

	for _, e := range numbers {
		sum += e

	}

	ch <- float64(sum) / float64(len(numbers))
}

func getMax(numbers []int, ch chan int) {
	var max = numbers[0]

	for _, e := range numbers {
		if max < e {
			max = e
		}
	}
	ch <- max
}

func sendMessage(ch chan<- string) {
	for i := 0; i < 20; i++ {
		ch <- fmt.Sprintf("data %d", i)
	}
	close(ch)
}

func printMessage(ch <-chan string) {
	for message := range ch {
		fmt.Println(message)
	}
}
