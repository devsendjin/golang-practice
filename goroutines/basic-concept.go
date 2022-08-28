package goroutines

import (
	"fmt"
	"os"
)

func BasicConcept() {
	greet()
	storeData("This is some dummy data!", "dummy-data.txt")

	channel1 := make(chan int)
	channel2 := make(chan int)

	go storeMoreData(50000, "50000_1.txt", channel1)
	go storeMoreData(50000, "50000_2.txt", channel2)

	<-channel1
	<-channel2
}

func greet() {
	fmt.Println("Hi there!")
}

func storeData(storableText string, fileName string) {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)

	if err != nil {
		fmt.Println("Creating the file failed. Exiting.")
		return
	}

	defer file.Close()

	_, err = file.WriteString(storableText)

	if err != nil {
		fmt.Println("Writing to the file failed.")
	}
}

func storeMoreData(lines int, fileName string, c chan int) {
	for i := 0; i < lines; i++ {
		text := fmt.Sprintf("Line %v - Dummy Data\n", i)
		storeData(text, fileName)
	}

	fmt.Printf("-- Done storing %v lines of text -- %v\n", lines, fileName)

	c <- 1
}
