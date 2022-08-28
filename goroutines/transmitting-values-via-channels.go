package goroutines

import (
	"fmt"
	"math/rand"
	"time"
)

var source = rand.NewSource(time.Now().Unix())
var randN = rand.New(source)

func ConcurrentlySelectWithBufferedChannels() {
	fmt.Println("ConcurrentlySelectWithBufferedChannels")

	const (
		LIMITER          = 3
		GOROUTINE_AMOUNT = 2
	)

	limiter := make(chan int, LIMITER)
	x := make(chan int)
	y := make(chan int)

	// only <LIMITER> func calls in parralel is possible
	go generateValueConcurrentlyWithBufferedChannels(x, limiter)
	go generateValueConcurrentlyWithBufferedChannels(y, limiter)

	select {
	case a := <-x:
		fmt.Printf("X finished first, value is %v\n", a)
	case b := <-y:
		fmt.Printf("Y finished first, value is %v\n", b)
	}
}

func ConcurrentlyWithBufferedChannels() {
	fmt.Println("TransmittingConcurrently")

	const (
		LIMITER          = 3
		GOROUTINE_AMOUNT = 4
	)

	c := make(chan int)
	limiter := make(chan int, LIMITER)

	// only <LIMITER> func calls in parralel is possible
	go generateValueConcurrentlyWithBufferedChannels(c, limiter)
	go generateValueConcurrentlyWithBufferedChannels(c, limiter)
	go generateValueConcurrentlyWithBufferedChannels(c, limiter)
	go generateValueConcurrentlyWithBufferedChannels(c, limiter)

	sum := 0
	i := 0

	for num := range c {
		fmt.Println("Num:", num)
		sum += num
		i++
		if i == GOROUTINE_AMOUNT { // <GOROUTINE_AMOUNT> - amount of goroutines
			close(c)
		}
	}

	fmt.Println("Sum:", sum)
}

func generateValueConcurrentlyWithBufferedChannels(c chan int, limit chan int) {
	limit <- 1

	sleepTime := randN.Intn(6)
	fmt.Println("WithBufferedChannels sleepTime", sleepTime)
	time.Sleep(time.Duration(sleepTime) * time.Second)

	c <- randN.Intn(10)

	<-limit
}

func ConcurrentlyWithUnbufferedChannels() {
	fmt.Println("TransmittingConcurrently")

	c := make(chan int)

	go generateValueConcurrentlyWithUnbufferedChannels(c)
	go generateValueConcurrentlyWithUnbufferedChannels(c)
	go generateValueConcurrentlyWithUnbufferedChannels(c)
	go generateValueConcurrentlyWithUnbufferedChannels(c)

	sum := 0
	i := 0

	for num := range c {
		fmt.Println("Num:", num)
		sum += num
		i++
		if i == 4 { // 4 - amount of goroutines
			close(c)
		}
	}

	fmt.Println("Sum:", sum)
}

func generateValueConcurrentlyWithUnbufferedChannels(c chan int) {
	sleepTime := randN.Intn(6)
	fmt.Println("WithUnbufferedChannels sleepTime", sleepTime)
	time.Sleep(time.Duration(sleepTime) * time.Second)

	c <- randN.Intn(10)
}

func TransmittingSequentially() {
	fmt.Println("TransmittingSequentially")

	x := generateValueSequentially()
	y := generateValueSequentially()

	fmt.Println("x, y", x, y)

	sum := x + y

	fmt.Println("Sum:", sum)
}

func generateValueSequentially() int {
	sleepTime := randN.Intn(6)
	fmt.Println("sleepTime", sleepTime)
	time.Sleep(time.Duration(sleepTime) * time.Second)

	return randN.Intn(10)
}
