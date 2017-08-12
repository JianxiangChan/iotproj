package main

import (
	"fmt"
	"hug/applog"
	"hug/backend"
	"hug/config"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/codegangsta/cli"
	log "github.com/fatih/color"
)

const (
	UPLINK_TOPIC   = "node/+/uplink"
	DOWNLINK_TOPIC = "node/+/downlink"
)

func run(c *cli.Context) error {
	log.Yellow("============监控MQTT所有通信的上下行消息============")
	conf, err := config.ReadConfig(c.String("conf"))
	if err != nil {
		fmt.Println("read from conf fail!", c.String("conf"))
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

	go submq(conf, UPLINK_TOPIC, 0)

	go submq(conf, DOWNLINK_TOPIC, 1)

	//quit when receive end signal
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	fmt.Printf("signal received signal %v", <-sigChan)
	fmt.Println("shutting down server")
	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "sug"
	app.Usage = "sug"
	app.Copyright = "sub@gmail.com"
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

func submq(conf config.Config, topic string, color int) {
	mqclient, err := backend.NewBackend(conf.Server, conf.Username, conf.Password)
	fmt.Printf("mqclient = [%v], topic = [%v]\n", mqclient, topic)
	fmt.Printf("err = [%v]\n", err)
	if mqclient == nil {
		return
	}
	mqclient.SubscribeTopic(topic)
	for rx := range mqclient.RxDataChan() {
		if color == 0 {
			log.Green("time = %v rxdata = %v", time.Now(), rx)
		} else {
			log.Cyan("time = %v rxdata = %v", time.Now(), rx)
		}

	}
}
