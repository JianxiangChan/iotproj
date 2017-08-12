package config

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	"gopkg.in/ini.v1"
)

type Config struct {
	Server   string `ini:"server"`
	Username string `ini:"username"`
	Password string `ini:"password"`
	//Topic
	Topic string `ini:"topic"`
	Mac   string `ini:"mac"`
	//LOG
	LogToFile   bool   `ini:"log_to_file"`
	LogDirWin   string `ini:"log_dir_win"`
	LogDirLinux string `ini:"log_dir_linux"`
	LogPrefix   string `ini:"log_prefix"`
}

func (c Config) String() string {
	mqtt := fmt.Sprintf("MQTT:[%v]/[%v]/[%v]", c.Server, c.Username, c.Password)
	topic := fmt.Sprintf("[Topic: %v] [Mac: %v]", c.Topic, c.Mac)
	log := fmt.Sprintf("LOG:[ToFile:%v][win:%v]/[linux:%v]:[prefix:%v]", c.LogToFile, c.LogDirWin, c.LogDirLinux, c.LogPrefix)

	return mqtt + ", " + topic + ", " + log
}

//Read Server's Config Value from "path"
func ReadConfig(path string) (Config, error) {
	var config Config
	conf, err := ini.Load(path)
	if err != nil {
		log.Println("load config file fail!")
		return config, err
	}
	conf.BlockMode = false
	err = conf.MapTo(&config)
	if err != nil {
		log.Println("mapto config file fail!")
		return config, err
	}
	return config, nil
}
