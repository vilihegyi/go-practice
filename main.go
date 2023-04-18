package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, Go!")
	fmt.Println("Going on forever now...")

	go forever()
	select {} // block forever
}

func forever() {
	for {
		fmt.Printf("%v+\n", time.Now())
		time.Sleep(time.Second)
	}
}
