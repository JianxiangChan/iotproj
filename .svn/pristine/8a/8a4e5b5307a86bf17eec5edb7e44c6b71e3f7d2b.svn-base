package hardware

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

//输入输出模式
const (
	Mode_INPUT  string = "in"  //输入模式
	Mode_OUTPUT string = "out" //输出模式
)

//引脚定义
const (
	LED_PIN int = 26
)

//输出值定义
const (
	LED_HIGH int = 1
	LED_LOW  int = 0
)

const (
	GPIO_EXPORT    = "/sys/class/gpio/export"
	GPIO_DIRECTION = "/sys/class/gpio/gpio%v/direction"
	GPIO_VALUE     = "/sys/class/gpio/gpio%v/value"
	GPIO_UNEXPOET  = "/sys/class/gpio/unexport"
)

//定义GPIO端口
type GPIODriver struct {
	pin      int
	pinMode  string
	pinValue int
}

func NewGPIODriver(ledpin int, mode string, defvalue int) *GPIODriver {
	driver := &GPIODriver{
		pin:      ledpin,
		pinMode:  mode,
		pinValue: defvalue,
	}
	return driver
}

//初始化
func (dr *GPIODriver) Init() {
	log.Info("GPIODriver Init")
	//EXPORT
	if err := dr.writeValue(GPIO_EXPORT, dr.pin); err != nil {
		log.Info("EXPORT Fail:", err)
	}
	//Direction
	direction_str := fmt.Sprintf(GPIO_DIRECTION, dr.pin)
	log.Info("direction_str = ", direction_str)
	if err := dr.writeString(direction_str, dr.pinMode); err != nil {
		log.Info("DIRECTION Fail:", err)
	}
}

//打开LED
func (dr *GPIODriver) Open() {
	log.Info("GPIODriver Open")
	dr.pinValue = LED_HIGH
	valuestr := fmt.Sprintf(GPIO_VALUE, dr.pin)
	log.Info("valuestr = ", valuestr)
	if err := dr.writeValue(valuestr, LED_HIGH); err != nil {
		log.Info("Open Fail:", err)
	}
}

//关闭LED
func (dr *GPIODriver) Close() {
	log.Info("GPIODriver Close")
	dr.pinValue = LED_LOW
	valuestr := fmt.Sprintf(GPIO_VALUE, dr.pin)
	fmt.Println("valuestr = ", valuestr)
	if err := dr.writeValue(valuestr, LED_LOW); err != nil {
		log.Info("Close Fail:", err)
	}
}

//func (dr *GPIODriver) Export() error {
//	fmt.Println("GPIODriver Export")

//}

//func (dr *GPIODriver) UnExport() error {
//	fmt.Println("GPIODriver UnExport")
//}

//func (dr *GPIODriver) UnDirection() error {
//	fmt.Println("GPIODriver UnDirection")
//}

//func (dr *GPIODriver) Read() error {
//	fmt.Println("GPIODriver Read")
//}

//func (dr *GPIODriver) Write() error {
//	fmt.Println("GPIODriver Write")
//}

func (dr *GPIODriver) readValue(path string) (int, error) {
	file, err := os.OpenFile(path, os.O_RDWR, 0666)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return 0, err
	}
	str := string(bytes)
	log.Info("readValue = ", str)
	return strconv.Atoi(str)
}

func (dr *GPIODriver) writeValue(path string, value int) error {
	str := strconv.Itoa(value)
	log.Info("writeValue = ", str)
	file, err := os.OpenFile(path, os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(str)
	if err != nil {
		return err
	}
	return nil
}

func (dr *GPIODriver) writeString(path string, value string) error {
	log.Info("writeValue = ", value)
	file, err := os.OpenFile(path, os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(value)
	if err != nil {
		return err
	}
	return nil
}
