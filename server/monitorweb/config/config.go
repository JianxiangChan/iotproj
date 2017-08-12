package config

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	"gopkg.in/ini.v1"
)

type Config struct {
	//MQTT
	MqttServer   string `ini:"mqtt_server"`
	MqttUsername string `ini:"mqtt_username"`
	MqttPassword string `ini:"mqtt_password"`
	MqttSubTopic string `ini:"mqtt_subtopic"`
	MqttPubTopic string `ini:"mqtt_pubtopic"`
	//LOG
	LogDirWin   string `ini:"log_dir_win"`
	LogDirLinux string `ini:"log_dir_linux"`
	LogPrefix   string `ini:"log_prefix"`
	LogToFile   bool   `ini:"log_tofile"`
	//MYSQL
	Mysql string `ini:"mysql"`
	//HTTP
	HttpServerWin   string `ini:"http_server_win"`
	HttpServerLinux string `ini:"http_server_linux"`
}

func (c Config) String() string {
	mqtt := fmt.Sprintf("MQTT:[%v]/[%v]/[%v]/[Sub:%v]/[Pub:%v]", c.MqttServer, c.MqttUsername, c.MqttPassword, c.MqttSubTopic, c.MqttPubTopic)

	log := fmt.Sprintf("LOG:[win:%v]/[linux:%v]:[prefix:%v]:[tofile:%v]", c.LogDirWin, c.LogDirLinux, c.LogPrefix, c.LogToFile)

	mysql := fmt.Sprintf("MYSQL:[%v]", c.Mysql)
	return mqtt + ", " + log + ", " + mysql
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
