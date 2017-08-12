package applog

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/robfig/cron"
)

const (
	SPEC_LOG = "@daily"
)

//Auto Daily Save Log Manager
type AutoDailyLoger struct {
	dir    string
	prefix string
	file   *os.File
	cron   *cron.Cron
}

func NewAutoDailyLoger(dir string, prefix string) *AutoDailyLoger {
	c := cron.New()

	//init output
	name := fmt.Sprintf("%v.log", filepath.Join(dir, prefix+time.Now().Format("20060102")))
	fmt.Println("dir = ", dir, " ,name = ", name)
	os.MkdirAll(dir, 0777)
	file, _ := os.OpenFile(name, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0)
	if file != nil {
		log.SetOutput(file)
	}
	log.Info("NewAutoDailyLoger OK!!!")
	s := &AutoDailyLoger{
		dir:    dir,
		prefix: prefix,
		cron:   c,
		file:   file,
	}
	c.AddFunc(SPEC_LOG, s.changeLogerFile)
	return s
}

func (s *AutoDailyLoger) Start() {
	s.cron.Start()
	log.Info("AutoDailyLoger start")
}

func (s *AutoDailyLoger) Stop() {
	s.cron.Stop()
	if s.file != nil {
		s.file.Close()
		s.file = nil
	}
	log.Info("AutoDailyLoger stop")
}

func (s *AutoDailyLoger) changeLogerFile() {
	if s.file != nil {
		s.file.Close()
		s.file = nil
	}
	name := fmt.Sprintf("%v.log", filepath.Join(s.dir, s.prefix+time.Now().Format("20060102")))
	os.MkdirAll(s.dir, 0777)
	s.file, _ = os.OpenFile(name, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0)
	if s.file != nil {
		log.SetOutput(s.file)
	}
	log.Info("changeLogerFile OK!!!")
}
