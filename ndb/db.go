package ndb

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/zjutjh/jhgo/nlog"
)

// New 以指定配置创建实例
func New(conf Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.Username, conf.Password, conf.Host, conf.Port, conf.Database)
	var l logger.Interface
	if conf.OpenLogger {
		l = newDBLogger(nlog.Pick(), logger.Config{
			SlowThreshold:             conf.SlowThreshold,
			Colorful:                  conf.Colorful,
			IgnoreRecordNotFoundError: conf.IgnoreRecordNotFoundError,
			ParameterizedQueries:      conf.ParameterizedQueries,
			LogLevel:                  conf.LogLevel,
		})
	}
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{
		Logger: l,
	})
	if err != nil {
		return nil, fmt.Errorf("创建gorm实例错误: %w", err)
	}
	sd, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("获取gorm中的DB实例错误: %w", err)
	}
	sd.SetMaxIdleConns(conf.MaxIdleConns)
	sd.SetMaxOpenConns(conf.MaxOpenConns)
	sd.SetConnMaxLifetime(conf.ConnMaxLifetime)
	sd.SetConnMaxIdleTime(conf.ConnMaxIdleTime)

	return db, nil
}
