package main

import (
	"fmt"
	"monitor/applog"
	"monitor/backend"
	"monitor/config"
	"monitor/dataserver"
	"runtime"

	"os"
	"os/signal"
	"syscall"

	log "github.com/Sirupsen/logrus"

	"github.com/codegangsta/cli"
)

func run(c *cli.Context) error {
	conf, err := config.ReadConfig(c.String("conf"))
	if err != nil {
		log.Error("read from conf fail!", c.String("conf"))
		return err
	}
	fmt.Println("conf =  ", conf)

	fmt.Println("runtime.GOOS = ", runtime.GOOS)

	var logger *applog.AutoDailyLoger
	if runtime.GOOS == "windows" {
		logger = applog.NewAutoDailyLoger(conf.LogDirWin, conf.LogPrefix)
	} else {
		logger = applog.NewAutoDailyLoger(conf.LogDirLinux, conf.LogPrefix)
	}
	logger.Start()
	defer logger.Stop()

	//Create MQTT Backend
	mqtt, err := backend.NewBackend(conf.MqttServer, conf.MqttUsername, conf.MqttPassword, conf.MqttPubTopic)
	if err != nil {
		log.Error("can not connect mqtt server")
		return err
	}
	defer mqtt.Close()

	// Subscribe Topic
	topic := conf.MqttSubTopic
	err = mqtt.SubscribeTopic(topic)
	if err != nil {
		log.Errorf("SubscribeTopic %v Error", topic)
		return err
	}
	defer mqtt.UnSubscribeTopic(topic)

	//start DataServer
	ds := dataserver.NewDataServer(mqtt, conf.Mac)
	go func() {
		ds.Start()
	}()
	defer ds.Stop()

	//start http server
	//	go func() {
	//		if runtime.GOOS == "windows" {
	//			httpserver.StartHttpServer(mqtt, db, conf.HttpServerWin)
	//		} else {
	//			httpserver.StartHttpServer(mqtt, db, conf.HttpServerLinux)
	//		}
	//	}()

	//quit when receive end signal
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	log.Infof("signal received signal %v", <-sigChan)
	log.Warn("shutting down server")
	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "monitor"
	app.Usage = "monitor pi board and transmit temprature and control led from remote server!"
	app.Copyright = "monitor123456@gmail.com"
	app.Version = "0.0.3"
	app.Action = run
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "conf,c",
			Usage:  "Set conf path here",
			Value:  "appserver.conf",
			EnvVar: "APP_CONF",
		},
	}
	app.Run(os.Args)
}
