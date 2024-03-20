package connect

import (
	"context"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"net/http"
	"time"
	"work.ctyun.cn/git/GoStack/gostone/execption"
)

var db *gorm.DB
var err error

func InitDB() {
	db, err = getNewConnect()
	if err != nil {
		panic(err)
	}

}

func getNewConnect() (*gorm.DB, error) {
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN: AppConf.GoStone.Database.Url,
	}), &gorm.Config{
		PrepareStmt: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: NewLogrusLogger(),
	})
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(AppConf.GoStone.MaxIdleConns)
	sqlDB.SetMaxOpenConns(AppConf.GoStone.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Minute)
	return db, nil
}

func GetMysqlConnect() *gorm.DB {
	if AppConf.GoStone.Database.Timeout == 0 {
		AppConf.GoStone.Database.Timeout = 5
	}
	getChannel := make(chan *gorm.DB)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(AppConf.GoStone.Database.Timeout)*time.Second)
	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Error("get mysql connect fail please check mysql status")
				return
			default:
				if db != nil {
					sqlDB, err := db.DB()
					if err == nil {
						if err = sqlDB.Ping(); err == nil {
							getChannel <- db
							return
						}
					}
				}
				db, err = getNewConnect()
				if err == nil {
					getChannel <- db
					return
				}
			}
		}
	}()
	defer cancel()
	select {
	case mysqlDb := <-getChannel:
		return mysqlDb
	case <-time.After(time.Duration(AppConf.GoStone.Database.Timeout) * time.Second):
		log.Error("get mysql connect fail please check mysql status")
		panic(execption.GoStoneError{
			Message: "get mysql connect fail please check mysql status",
			Code:    http.StatusInternalServerError,
		})
	}
}
