package database

import (
	"monitorweb/model"

	log "github.com/Sirupsen/logrus"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

//DBEngin
type DBEngine struct {
	dbengine *xorm.Engine
}

//Create DBEngin
func NewDBEngine(dburl string) (*DBEngine, error) {
	db, err := xorm.NewEngine("mysql", dburl)
	if err != nil {
		return nil, err
	}

	db.ShowSQL(false)
	log.Infof("connect to database(%v) server OK!", dburl)

	db.Sync2(new(model.TabAllSensorData), new(model.TabAllSensorDataLast))

	engine := DBEngine{
		dbengine: db,
	}

	return &engine, nil
}

//Insert All Sensor Data
func (db *DBEngine) InsertTabAllSensorData(rx model.TabAllSensorData) error {
	_, err := db.dbengine.Insert(&rx)
	if err != nil {
		log.Errorf("InsertTabAllSensorData  fail %v!", err)
	}
	return err
}

//Query All Sensor last one data
func (db *DBEngine) QueryAllLastData() ([]model.TabAllSensorDataLast, error) {
	all := make([]model.TabAllSensorDataLast, 0)
	err := db.dbengine.Desc("id").Find(&all)
	return all, err
}

//Query last 100 data by mac
func (db *DBEngine) QueryDetailData(mac string) ([]model.TabAllSensorData, error) {
	all := make([]model.TabAllSensorData, 0)
	err := db.dbengine.Where("mac = ?", mac).Desc("id").Limit(100).Find(&all)
	return all, err
}
