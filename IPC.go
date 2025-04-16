// Name: Neel Ranawat
// KSUID: 000988101
// Class: CS 4308 W03 Spring 2025
// This program demonstrates basic IPC with Go
package main

import (
	"fmt"
	"time"
)

// serves spaghetti
func Server(ch chan<- int) {
	for i := 0; i < 3; i++ {
		ch <- i
		fmt.Printf("Spaghetti cooked and served: %d\n", i)
		time.Sleep(time.Second) // Simulates serving
	}
	close(ch) // Closes channel
}

func ChinesePhilosopher(ch <-chan int) {
	for num := range ch {
		fmt.Println("Chinese philosopher is philosophizing")
		time.Sleep(500 * time.Millisecond) // Simulates thinking
		fmt.Printf("Spaghetti eaten: %d\n", num)
	}
}

func main() {
	//starts a channel for two
	ch := make(chan int, 2)

	go Server(ch)          //Starts serving in channel
	ChinesePhilosopher(ch) //Starts thinking and eating inchannel

	fmt.Println("The Chinese Philosopher figured it out!")
}
