package backend

import (
	"encoding/json"
	"fmt"
	"monitorweb/packets"
	"sync"

	log "github.com/Sirupsen/logrus"
	"github.com/eclipse/paho.mqtt.golang"
)

// Backend implements a MQTT pub-sub backend.
type Backend struct {
	conn     mqtt.Client
	rxChan   chan packets.TempPacket //接收设备上行主题
	mutex    sync.RWMutex
	pubtopic string //发布主题 node/%s/downlink
	subtopic string //订阅主题
}

// NewBackend creates a new Backend.
func NewBackend(server, username, password string, pubtopic string) (*Backend, error) {
	b := Backend{
		rxChan:   make(chan packets.TempPacket),
		pubtopic: pubtopic,
		subtopic: "",
	}

	opts := mqtt.NewClientOptions()
	opts.AddBroker(server)
	opts.SetUsername(username)
	opts.SetPassword(password)
	opts.SetOnConnectHandler(b.onConnected)
	opts.SetConnectionLostHandler(b.onConnectionLost)

	log.Infof("backend/mqttpubsub: connecting to mqtt broker %v", server)
	b.conn = mqtt.NewClient(opts)
	if token := b.conn.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}

	return &b, nil
}

// Close closes the backend.
func (b *Backend) Close() {
	b.conn.Disconnect(250) // wait 250 milisec to complete pending actions
	log.Info("backend/mqttpubsub: Disconnect mqtt broker")
}

// RxDataChan returns the TabRxData channel.
func (b *Backend) RxDataChan() chan packets.TempPacket {
	return b.rxChan
}

// Subscribe RxData
func (b *Backend) SubscribeTopic(topic string) error {
	defer b.mutex.Unlock()
	b.mutex.Lock()

	log.Infof("backend/mqttpubsub: subscribing to topic %v", topic)
	if token := b.conn.Subscribe(topic, 2, b.rxDataHandler); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	b.subtopic = topic
	return nil
}

// UnSubscribe RxData
func (b *Backend) UnSubscribeTopic(topic string) error {
	defer b.mutex.Unlock()
	b.mutex.Lock()

	log.Infof("backend/mqttpubsub: unsubscribing from topic %v", topic)
	if token := b.conn.Unsubscribe(topic); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	b.subtopic = ""
	return nil
}

//下行控制LED帧
func (b *Backend) PublishPacket(pkt packets.LedPacket) error {
	bytes, err := json.Marshal(pkt)
	if err != nil {
		return err
	}
	topic := fmt.Sprintf(b.pubtopic, pkt.Mac)

	log.WithFields(log.Fields{
		"topic": topic,
		"bytes": bytes,
	}).Info("publish:")
	if token := b.conn.Publish(topic, 2, false, bytes); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}

func (b *Backend) rxDataHandler(c mqtt.Client, msg mqtt.Message) {
	log.Infof("backend/mqttpubsub: packet received from topic %v", msg.Topic())
	var rxData packets.TempPacket
	if err := json.Unmarshal(msg.Payload(), &rxData); err != nil {
		log.Errorf("backend/mqttpubsub: decode rxData error: %s", err)
		return
	}
	log.Infof("backend/mqttpubsub: packet received  %v", rxData)
	b.rxChan <- rxData
}

func (b *Backend) onConnected(c mqtt.Client) {
	defer b.mutex.RUnlock()
	b.mutex.RLock()
	if b.subtopic == "" {
		return
	}
	log.Info("backend/mqttpubsub: onConnected to mqtt broker")
	if token := b.conn.Subscribe(b.subtopic, 2, b.rxDataHandler); token.Wait() && token.Error() != nil {
		return
	}
}

func (b *Backend) onConnectionLost(c mqtt.Client, reason error) {
	log.Errorf("backend/mqttpubsub: mqtt onConnectionLost error: %s", reason)
}
