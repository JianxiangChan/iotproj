package hardware

import (
	"fmt"
)

/*
#cgo CFLAGS: -std=c99 -I ./
#cgo LDFLAGS: -L.  -lwiringPi
#include <stdio.h>
#include <stdlib.h>
#include <wiringPi.h>
*/
import "C"

//import "unsafe"

//输入输出模式
const (
	Mode_INPUT  int = 0 //输入模式
	Mode_OUTPUT int = 1 //输出模式
)

//引脚定义
const (
	LED_PIN int = 25
)

//输出值定义
const (
	LED_OPEN  int = 1
	LED_CLOSE int = 0
)

type LEDDriver struct {
	pin      int
	pinMode  int
	pinValue int
}

func NewLEDDriver() *LEDDriver {
	driver := &LEDDriver{
		pin:      LED_PIN,
		pinMode:  Mode_OUTPUT,
		pinValue: LED_CLOSE,
	}
	return driver
}

//初始化
func (driver *LEDDriver) Init() {
	fmt.Println("LEDDriver Init")
	//初始化板子
	C.wiringPiSetup()
	//设置为输出模式
	C.pinMode(C.int(driver.pin), C.int(driver.pinMode))
	C.digitalWrite(C.int(driver.pin), C.int(driver.pinValue))
}

//打开LED
func (driver *LEDDriver) Open() {
	fmt.Println("LEDDriver Open")
	C.digitalWrite(C.int(driver.pin), C.int(LED_OPEN))
	driver.pinValue = LED_OPEN
}

//关闭LED
func (driver *LEDDriver) Close() {
	fmt.Println("LEDDriver Close")
	C.digitalWrite(C.int(driver.pin), C.int(LED_CLOSE))
	driver.pinValue = LED_CLOSE
}
