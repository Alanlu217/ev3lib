package main

import (
	"github.com/Alanlu217/ev3lib/ev3lib"
	"github.com/Alanlu217/ev3lib/tests/configs/bopConfig"
)

func main() {
	config := bopConfig.NewBopConfig()

	ev3lib.RunCommand(config.GetCommandPages().Pages[0].Commands[1].Command)
}
