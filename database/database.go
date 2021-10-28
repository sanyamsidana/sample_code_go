package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"sync"
)
var db *gorm.DB
var dberr error
var once sync.Once

type DBConfig struct {
	DBHOST, DBUSER, DBNAME, DBPASSWORD, DBPORT string
}

func GetDbInstance() (*gorm.DB, error) {
	once.Do(func() {
		db, dberr = initDatabase()
	})
	if db == nil {
		return nil, dberr
	}
	err := db.DB().Ping()
	if err != nil {
		return nil, dberr
	}
	return db, nil

}

func initDatabase() (*gorm.DB, error) {
	dbconfig, err := GetDBConfig()
	if err  != nil {
		return nil, err
	}
	return newDbConnection(dbconfig)
}

func newDbConnection(config *DBConfig) (*gorm.DB, error) {
	Db, Err := gorm.Open(
		"postgres",
		"host="+config.DBHOST+" user="+config.DBUSER+" port="+config.DBPORT+
			" dbname="+config.DBNAME+" sslmode=disable password="+config.DBPASSWORD)
	if Err != nil {
		return nil, Err
	}
	return Db, nil
}

func GetDBConfig() (*DBConfig, error) {
	return &DBConfig{
		DBHOST:           "assess-clusterid-temp.cluster-c8q7plqsqewk.us-east-2.rds.amazonaws.com",
		DBNAME:           "assess_eventbus",
		DBUSER:           "eventbus",
		DBPASSWORD:       "jw8s0F4eventbus",
		DBPORT:           "5432",
	}, nil
}
func Hello(){
	fmt.Println("hello-2")

}