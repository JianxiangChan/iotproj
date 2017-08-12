package main

import (
	//"fmt"
	"led2/hardware"
	"time"
)

func main() {
	leddr := hardware.NewGPIODriver(hardware.LED_PIN, hardware.Mode_OUTPUT, hardware.LED_LOW)
	leddr.Init()
	for {
		leddr.Open()
		time.Sleep(2 * time.Second)
		leddr.Close()
		time.Sleep(2 * time.Second)
	}
}
