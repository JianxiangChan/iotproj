package dataserver

import (
	"monitorweb/backend"
	"monitorweb/database"
	"monitorweb/model"
	"monitorweb/packets"
	"time"

	log "github.com/Sirupsen/logrus"
)

type DataServer struct {
	mqtt *backend.Backend
	db   *database.DBEngine
}

func NewDataServer(mqtt *backend.Backend, db *database.DBEngine) *DataServer {
	server := &DataServer{
		mqtt: mqtt,
		db:   db,
	}
	return server
}

func (s *DataServer) Start() {
	go func() {
		for rxPayload := range s.mqtt.RxDataChan() {
			s.handleTempPacket(rxPayload)
		}
	}()
	log.Info("DataServer start")
}

func (s *DataServer) Stop() {
	log.Info("DataServer stop")
}

//处理上行状态数据帧
func (s *DataServer) handleTempPacket(pkt packets.TempPacket) {
	log.Info("handleTempPacket Inssert to Database")
	var data model.TabAllSensorData
	data.Mac = pkt.Mac
	data.Led = pkt.LED
	data.Temperature = pkt.Temperature
	data.CreateDate = time.Now()
	s.db.InsertTabAllSensorData(data)
}
