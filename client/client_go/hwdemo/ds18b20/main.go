package main

import (
	"ds18b20/hardware"
	"fmt"
	"time"
)

func main() {
	fmt.Println("=============DS18b20 Golang Demo =============")

	ds18b20 := hardware.NewDS18b20Driver()
	ds18b20.Init()
	for {
		temp := ds18b20.Read()
		fmt.Println("temp = ", temp)
		time.Sleep(5 * time.Second)

	}
}
