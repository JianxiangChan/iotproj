package main

import (
	"fmt"
	"led/hardware"
	"time"
)

func main() {
	fmt.Println("=============LED Golang Demo =============")
	leddr := hardware.NewLEDDriver()
	leddr.Init()
	for {
		leddr.Open()
		time.Sleep(2 * time.Second)
		leddr.Close()
		time.Sleep(2 * time.Second)
	}
}
