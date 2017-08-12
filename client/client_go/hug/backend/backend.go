package backend

import (
	"encoding/json"
	//"fmt"
	"sync"

	log "github.com/Sirupsen/logrus"

	"github.com/eclipse/paho.mqtt.golang"
)

// Backend implements a MQTT pub-sub backend.
type Backend struct {
	conn   mqtt.Client
	rxChan chan string
	mutex  sync.RWMutex
	topic  string
}

// NewBackend creates a new Backend.
func NewBackend(server, username, password string) (*Backend, error) {
	b := Backend{
		rxChan: make(chan string),
		topic:  "",
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
func (b *Backend) RxDataChan() chan string {
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
	b.topic = topic
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
	b.topic = ""
	return nil
}

func (b *Backend) Publish(topic string, v interface{}) error {
	bytes, err := json.Marshal(v)
	if err != nil {
		return err
	}

	//	log.WithFields(log.Fields{
	//		"topic": topic,
	//	}).Info("publish:")
	if token := b.conn.Publish(topic, 2, false, bytes); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}

func (b *Backend) rxDataHandler(c mqtt.Client, msg mqtt.Message) {
	//log.Infof("backend/mqttpubsub: packet received from topic %v", msg.Topic())
	rxData := msg.Topic() + " <----> " + string(msg.Payload())
	b.rxChan <- rxData
}

func (b *Backend) onConnected(c mqtt.Client) {
	defer b.mutex.RUnlock()
	b.mutex.RLock()
	if b.topic == "" {
		return
	}
	log.Info("backend/mqttpubsub: onConnected to mqtt broker")
	if token := b.conn.Subscribe(b.topic, 2, b.rxDataHandler); token.Wait() && token.Error() != nil {
		return
	}
}

func (b *Backend) onConnectionLost(c mqtt.Client, reason error) {
	log.Errorf("backend/mqttpubsub: mqtt onConnectionLost error: %s", reason)
}
