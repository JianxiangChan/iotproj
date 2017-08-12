package main

import (
	"fmt"
	"hug/applog"
	"hug/backend"
	"hug/config"
	"hug/nclient"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/codegangsta/cli"
	log "github.com/fatih/color"
)

func run(c *cli.Context) error {
	log.Yellow("============模拟下行控制对应的树莓派板子============")
	log.Yellow("=====每个树莓派通过它的Mac地址进行唯一标识区别=======")
	conf, err := config.ReadConfig(c.String("conf"))
	if err != nil {
		fmt.Printf("read from conf fail!", c.String("conf"))
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

	go mqttclient(conf)

	//quit when receive end signal
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	fmt.Printf("signal received signal %v", <-sigChan)
	fmt.Println("shutting down server")
	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "hug"
	app.Usage = "hug"
	app.Copyright = "hug@gmail.com"
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

func mqttclient(conf config.Config) {
	mqclient, err := backend.NewBackend(conf.Server, conf.Username, conf.Password)
	fmt.Printf("mqclient = [%v]\n", mqclient)
	fmt.Printf("err = [%v]\n", err)
	if mqclient == nil {
		return
	}
	ncli := nclient.NewNClient(mqclient, conf.Topic, conf.Mac)
	ncli.Start()
}
