package main

import (
	"fmt"
	"log"
	"time"

	"github.com/ev3go/ev3dev"
)

func main() {
	b := ev3dev.ButtonPoller{}

	t := time.NewTicker(time.Millisecond * 500)

	for {
		start := time.Now()
		val, err := b.Poll()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(val)

		fmt.Println(time.Since(start))

		<-t.C
	}
}
