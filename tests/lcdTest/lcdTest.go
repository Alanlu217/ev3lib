package main

import (
	"fmt"
	"time"

	"github.com/Alanlu217/ev3lib/ev3lib/ev3"
)

func main() {
	// l := ev3.LCD
	hub := ev3.NewEV3()
	t := time.NewTicker(time.Second)

	for {
		start := time.Now()
		hub.ClearScreen()

		hub.DrawText(0, 0, "Hello My Name is Alan")
		// fmt.Println("AKLSJDALSKDJLAKSJD")

		fmt.Println(time.Since(start))

		<-t.C
	}
}
