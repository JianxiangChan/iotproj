//go:generate go-bindata -prefix ./public/ -o=assets/public/public_gen.go -pkg=public public/...
//go:generate go-bindata -prefix ./templates/ -o=assets/templates/templates_gen.go -pkg=templates templates/...

package main

import (
	"fmt"
	"monitorweb/applog"
	"monitorweb/backend"
	"monitorweb/config"
	"monitorweb/database"
	"monitorweb/dataserver"
	"monitorweb/httpserver"
	"monitorweb/model"
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

	if conf.LogToFile {
		var logger *applog.AutoDailyLoger
		if runtime.GOOS == "windows" {
			logger = applog.NewAutoDailyLoger(conf.LogDirWin, conf.LogPrefix)
		} else {
			logger = applog.NewAutoDailyLoger(conf.LogDirLinux, conf.LogPrefix)
		}
		logger.Start()
		defer logger.Stop()
	}

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

	//Connect to Mysql Database
	db, err := database.NewDBEngine(conf.Mysql)
	if err != nil {
		log.Errorf("connect to mysql server fail %v!", err)
		return err
	}

	//start DataServer
	ds := dataserver.NewDataServer(mqtt, db)
	go func() {
		ds.Start()
	}()
	defer ds.Stop()

	mqttconf, err1 := model.ReadMqttConfig(c.String("conf"))
	if err1 != nil {
		log.Error("read from mqtt conf fail!", c.String("conf"))
		return err1
	}
	fmt.Println(mqttconf)

	//start http server
	go func() {
		if runtime.GOOS == "windows" {
			httpserver.StartHttpServer(mqtt, db, mqttconf, conf.HttpServerWin)
		} else {
			httpserver.StartHttpServer(mqtt, db, mqttconf, conf.HttpServerLinux)
		}
	}()

	//quit when receive end signal
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	log.Infof("signal received signal %v", <-sigChan)
	log.Warn("shutting down server")
	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "monitorweb"
	app.Usage = "monitorweb view  temprature and control led web server!"
	app.Copyright = "monitorweb123456@gmail.com"
	app.Version = "0.0.1"
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
