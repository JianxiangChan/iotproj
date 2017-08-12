package model

import (
	log "github.com/Sirupsen/logrus"
	"gopkg.in/ini.v1"
)

type MqttEntity struct {
	MqttTopic           string `ini:"mqttweb_topic"`
	MqttHostname        string `ini:"mqttweb_hostname"`
	MqttPort            int    `ini:"mqttweb_port"`
	MqttPath            string `ini:"mqttweb_path"`
	MqttUser            string `ini:"mqttweb_user"`
	MqttPass            string `ini:"mqttweb_pass"`
	MqttKeepAlive       int    `ini:"mqttweb_keepalive"`
	MqttTimeout         int    `ini:"mqttweb_timeout"`
	MqttSsl             bool   `ini:"mqttweb_ssl"`
	MqttCleanSession    bool   `ini:"mqttweb_cleanSession"`
	MqttLastWillTopic   string `ini:"mqttweb_lastWillTopic"`
	MqttLastWillQos     int    `ini:"mqttweb_lastWillQos"`
	MqttLastWillRetain  bool   `ini:"mqttweb_lastWillRetain"`
	MqttLastWillMessage string `ini:"mqttweb_lastWillMessage"`
}

func ReadMqttConfig(path string) (MqttEntity, error) {
	var mqtt MqttEntity
	mq, err := ini.Load(path)
	if err != nil {
		log.Println("load mqtt file fail!")
		return mqtt, err
	}
	mq.BlockMode = false
	err = mq.MapTo(&mqtt)
	if err != nil {
		log.Println("mapto config file fail!")
		return mqtt, err
	}

	return mqtt, nil
}
