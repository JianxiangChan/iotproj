package backend

import (
	//"encoding/base64"
	"encoding/json"
	//"errors"
	//"fmt"
	"monitor/packets"
	"sync"

	log "github.com/Sirupsen/logrus"

	"github.com/eclipse/paho.mqtt.golang"
	//"github.com/satori/go.uuid"
)

// Backend implements a MQTT pub-sub backend.
type Backend struct {
	conn     mqtt.Client
	rxChan   chan packets.LedPacket
	mutex    sync.RWMutex
	pubtopic string //发布主题
	subtopic string //订阅主题
}

// NewBackend creates a new Backend.
func NewBackend(server, username, password string, pubtopic string) (*Backend, error) {
	b := Backend{
		rxChan:   make(chan packets.LedPacket),
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
func (b *Backend) RxDataChan() chan packets.LedPacket {
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

func (b *Backend) PublishTempPacket(pkt packets.TempPacket) error {
	bytes, err := json.Marshal(pkt)
	if err != nil {
		return err
	}

	log.WithFields(log.Fields{
		"topic": b.pubtopic,
	}).Info("publish:")
	if token := b.conn.Publish(b.pubtopic, 2, false, bytes); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}

func (b *Backend) rxDataHandler(c mqtt.Client, msg mqtt.Message) {
	log.Infof("backend/mqttpubsub: packet received from topic %v", msg.Topic())
	var rxData packets.LedPacket
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
