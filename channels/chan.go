package main

import (
	"fmt"
	// "math/rand"
	"time"
)

// sending
func processNum(numChan chan int) {
	for num := range numChan {
		fmt.Println("Processing : ", num)
		time.Sleep(time.Second * 1)
	}
}

// receiving
func sum(result chan int, num1, num2 int) {
	numResult := num1 + num2
	result <- numResult
}

// go routine synchronization using channel
func task(done chan bool) {
	defer func() { done <- true }()
	fmt.Println("Task started")
	time.Sleep(time.Second * 2)
	fmt.Println("Task completed")
}

func emailSender(emailChan <-chan string, done chan bool) {
	defer func() { done <- true }()
	for email := range emailChan {
		fmt.Println("Sending email to : ", email)
		time.Sleep(time.Millisecond * 500)
	}
}

func channel() {

	// numChan := make(chan int)
	// go processNum(numChan)
	// for {
	// 	numChan <- rand.Intn(100)
	// }

	// messagechan := make(chan string)
	// messagechan <- "ping.." // blocking
	// msg := <-messagechan
	// fmt.Println(msg)

	// result := make(chan int)
	// go sum(result, 10, 20)
	// resultNum := <-result // blocking
	// fmt.Println("Result : ", resultNum)

	// done := make(chan bool)
	// go task(done)
	// <-done // blocking

	// emailChan := make(chan string, 5) // buffered channel with capacity of 5
	// done := make(chan bool)

	// go emailSender(emailChan, done)

	// for i := 0; i < 5; i++ {
	// 	emailChan <- fmt.Sprintf("%d@gmail.com", i)
	// }

	// fmt.Println("All emails sent")
	// close(emailChan) // this is important to close the channel after sending all emails
	// <-done

	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 42
	}()

	go func() {
		chan2 <- "pong!"
	}()

	for i := 0; i < 2; i++ {
		select {
		case num := <-chan1:
			fmt.Println("Received from chan1:", num)
		case msg := <-chan2:
			fmt.Println("Received from chan2:", msg)
		}
	}
}
