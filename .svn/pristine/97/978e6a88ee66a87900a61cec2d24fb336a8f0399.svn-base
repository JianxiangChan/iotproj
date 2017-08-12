package nclient

import (
	"fmt"
	"hug/backend"
	"time"

	log "github.com/fatih/color"
	"github.com/robfig/cron"
)

const (
	SPEC_MSG_INTEVAL string = "@every 30s"
)

type NClient struct {
	mqtt  *backend.Backend
	cron  *cron.Cron
	topic string
	mac   string
	count int //计数器，用户切换LED状态
}

func NewNClient(mqtt *backend.Backend, topic string, mac string) *NClient {
	c := cron.New()
	server := &NClient{
		mqtt:  mqtt,
		cron:  c,
		topic: topic,
		mac:   mac,
		count: 0,
	}
	c.AddFunc(SPEC_MSG_INTEVAL, server.handlePublish)
	return server
}

func (s *NClient) Start() {
	s.cron.Start()
	fmt.Println("NClient start")
}

func (s *NClient) Stop() {
	s.cron.Stop()
	fmt.Println("NClient stop")
}

func (s *NClient) handlePublish() {
	var pkt LedPacket
	pkt.Mac = s.mac
	if s.count%2 == 0 {
		pkt.LED = true
	} else {
		pkt.LED = false
	}
	s.count = s.count + 1
	if pkt.LED {
		log.Green("time = %v 远程打开LED", time.Now())
	} else {
		log.Green("time = %v 远程关闭LED", time.Now())
	}

	s.mqtt.Publish(s.topic, pkt)
}

//定义下行控制数据包
type LedPacket struct {
	Mac string `json:"mac"` //设备号，定义树莓派板唯一标识，使用MAC地址
	LED bool   `json:"led"` //控制LED灯亮 = true ,LED灯灭=flase
}
