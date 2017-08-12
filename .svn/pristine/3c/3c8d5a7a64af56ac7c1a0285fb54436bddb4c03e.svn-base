package hardware

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	log "github.com/Sirupsen/logrus"
)

const (
	PATH_HEAD string = "/sys/bus/w1/devices/"
	PATH_MID  string = "28-00000"
	PATH_TAIL string = "/w1_slave"
)

type DS18b20Driver struct {
	path string //温度传感器温度设备路径
}

func NewDS18b20Driver() *DS18b20Driver {
	driver := &DS18b20Driver{
		path: "",
	}
	return driver
}

//初始化
func (driver *DS18b20Driver) Init() {
	log.Info("DS18b20Driver Init")
	files, _ := ioutil.ReadDir(PATH_HEAD)
	midname := ""
	for _, f := range files {
		if strings.HasPrefix(f.Name(), PATH_MID) {
			midname = f.Name()
			break
		}
	}
	if midname != "" {
		driver.path = PATH_HEAD + midname + PATH_TAIL
		log.Info("driver.path = ", driver.path)
	} else {
		log.Info("DS18b20温度传感器不存在")
	}
}

//Read 温度
func (driver *DS18b20Driver) Read() float32 {
	log.Info("DS18b20Driver Read")
	if driver.path == "" {
		return 0.1
	}

	fi, err := os.Open(driver.path)
	if err != nil {
		log.Info("Open err = ", err)
		return 0.2
	}
	defer fi.Close()
	bytes, err := ioutil.ReadAll(fi)
	if err != nil {
		log.Info("Read err = ", err)
		return 0.3
	}
	content := string(bytes)
	log.Info(content)
	sp := strings.Split(content, "t=")
	if len(sp) != 2 {
		return 0.4
	}
	temp, err := strconv.Atoi(strings.TrimSpace(sp[1]))
	return float32(temp) * 1.0 / 1000.0
}
