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

		hub.DrawText(0, 19*0, "abcdefghijklmnop")
		// hub.DrawText(0, 19*1, "qrstuvwxyzABCDEF")
		// hub.DrawText(0, 19*2, "GHIJKLMNOPQRSTUV")
		// hub.DrawText(0, 19*3, "WXYZ0123456789")
		// hub.DrawText(0, 19*4, "#$%&'()*+,-./:;<")
		// hub.DrawText(0, 19*5, "=>?@[\\]^_`{|}~ ")
		// fmt.Println("AKLSJDALSKDJLAKSJD")

		fmt.Println(time.Since(start))

		<-t.C
	}
}
