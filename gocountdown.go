package main

import (
	"fmt"
	"time"
)

func main() {
	var count = 10
	for count > 0 {
		fmt.Println(count)
		time.Sleep(time.Second)
		count--
	}
	fmt.Println("Liftoff!")
             	var room string
	fmt.Println("enter a room: ")
	 fmt.Scanln(&room)
	 fmt.Println("good", room)
}
