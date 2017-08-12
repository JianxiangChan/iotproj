package dataserver

import (
	"monitor/backend"
	"monitor/hardware"
	"monitor/packets"

	log "github.com/Sirupsen/logrus"
	"github.com/robfig/cron"
)

//上报周期
const (
	SPEC_UP_LEVEL string = "@every 30s" //30s上报一次
)

type DataServer struct {
	mqtt    *backend.Backend
	cron    *cron.Cron
	leddr   *hardware.GPIODriver    //LED
	ds18b20 *hardware.DS18b20Driver //ds18b20
	mac     string                  //Pi 板子Mac地址
	temp    float32                 //Pi温度
	led     bool                    //LED关闭状态 = true le is open / = false led is close
}

func NewDataServer(mqtt *backend.Backend, mac string) *DataServer {
	c := cron.New()
	server := &DataServer{
		mqtt:    mqtt,
		cron:    c,
		leddr:   hardware.NewGPIODriver(hardware.LED_PIN, hardware.Mode_OUTPUT, hardware.LED_LOW),
		ds18b20: hardware.NewDS18b20Driver(),
		mac:     mac,
		temp:    0.0,
		led:     false,
	}
	server.leddr.Init()
	server.ds18b20.Init()
	c.AddFunc(SPEC_UP_LEVEL, server.handleTempPacket)
	return server
}

func (s *DataServer) Start() {
	go func() {
		for rxPayload := range s.mqtt.RxDataChan() {
			s.handleLedPacket(rxPayload)
		}
	}()
	s.cron.Start()
	log.Info("DataServer start")
}

func (s *DataServer) Stop() {
	s.cron.Stop()
	log.Info("DataServer stop")
}

//处理下行控制数据帧
func (s *DataServer) handleLedPacket(pkt packets.LedPacket) {
	log.Info("handleLedPacket LedPkt = ", pkt)
	s.led = pkt.LED
	if s.led {
		s.leddr.Open()
	} else {
		s.leddr.Close()
	}
	//处理下行数据后，及时反馈状态到服务
	s.handleTempPacket()
}

//处理上行状态数据帧
func (s *DataServer) handleTempPacket() {
	log.Info("handleTempPacket")
	//采集温度
	s.temp = s.ds18b20.Read()
	var tempPkt packets.TempPacket
	tempPkt.LED = s.led
	tempPkt.Mac = s.mac
	tempPkt.Temperature = s.temp
	err := s.mqtt.PublishTempPacket(tempPkt)
	if err != nil {
		log.Error("handleTempPacket error: ", err)
	}
}
