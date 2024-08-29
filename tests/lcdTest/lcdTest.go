package main

import (
	"github.com/Alanlu217/ev3lib/ev3lib"
	"github.com/Alanlu217/ev3lib/ev3lib/ev3"
	"time"
)

func main() {
	l := ev3.LCD
	hub := ev3.NewEV3()

	for {
		hub.ClearScreen()
		for i := 0; i < 128; i++ {
			for j := 0; j < 18; j++ {
				// Black
				l.Set(ev3lib.LCDPixelToIndex(10*j, i), 0)
				l.Set(ev3lib.LCDPixelToIndex(10*j, i)+1, 0)
				l.Set(ev3lib.LCDPixelToIndex(10*j, i)+2, 0)
				l.Set(ev3lib.LCDPixelToIndex(10*j, i)+3, 0)

				// White
				//l.Set(ev3lib.LCDPixelToIndex(11, i), 255)
				//l.Set(ev3lib.LCDPixelToIndex(11, i)+1, 255)
				//l.Set(ev3lib.LCDPixelToIndex(11, i)+2, 255)
				//l.Set(ev3lib.LCDPixelToIndex(11, i)+3, 255)

			}
		}
		for j := 0; j < 178; j++ {
			for i := 0; i < 13; i++ {
				// Black
				l.Set(ev3lib.LCDPixelToIndex(j, i*10), 0)
				l.Set(ev3lib.LCDPixelToIndex(j, i*10)+1, 0)
				l.Set(ev3lib.LCDPixelToIndex(j, i*10)+2, 0)
				l.Set(ev3lib.LCDPixelToIndex(j, i*10)+3, 0)

				// White
				//l.Set(ev3lib.LCDPixelToIndex(11, i), 255)
				//l.Set(ev3lib.LCDPixelToIndex(11, i)+1, 255)
				//l.Set(ev3lib.LCDPixelToIndex(11, i)+2, 255)
				//l.Set(ev3lib.LCDPixelToIndex(11, i)+3, 255)

			}
		}
		time.Sleep(time.Second)
	}
}
