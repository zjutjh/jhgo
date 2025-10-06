package ndb

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/zjutjh/mygo/nlog"
)

// New 以指定配置创建实例
func New(conf Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%s&loc=%s", conf.Username, conf.Password, conf.Host, conf.Port, conf.Database, conf.Charset, conf.ParseTime, conf.Loc)
	var l logger.Interface
	if conf.OpenLogger {
		l = newDBLogger(nlog.Pick(conf.Log), logger.Config{
			SlowThreshold:             conf.SlowThreshold,
			Colorful:                  conf.Colorful,
			IgnoreRecordNotFoundError: conf.IgnoreRecordNotFoundError,
			ParameterizedQueries:      conf.ParameterizedQueries,
			LogLevel:                  conf.LogLevel,
		})
	}
	db, err := gorm.Open(mysql.New(mysql.Config{
		DriverName:                    conf.DriverName,
		ServerVersion:                 conf.ServerVersion,
		DSN:                           dsn,
		SkipInitializeWithVersion:     conf.SkipInitializeWithVersion,
		DefaultStringSize:             conf.DefaultStringSize,
		DefaultDatetimePrecision:      &conf.DefaultDatetimePrecision,
		DisableWithReturning:          conf.DisableWithReturning,
		DisableDatetimePrecision:      conf.DisableDatetimePrecision,
		DontSupportRenameIndex:        conf.DontSupportRenameIndex,
		DontSupportRenameColumn:       conf.DontSupportRenameColumn,
		DontSupportForShareClause:     conf.DontSupportForShareClause,
		DontSupportNullAsDefaultValue: conf.DontSupportNullAsDefaultValue,
		DontSupportRenameColumnUnique: conf.DontSupportRenameColumnUnique,
	}), &gorm.Config{
		SkipDefaultTransaction:                   conf.SkipDefaultTransaction,
		FullSaveAssociations:                     conf.FullSaveAssociations,
		DryRun:                                   conf.DryRun,
		PrepareStmt:                              conf.PrepareStmt,
		DisableAutomaticPing:                     conf.DisableAutomaticPing,
		DisableForeignKeyConstraintWhenMigrating: conf.DisableForeignKeyConstraintWhenMigrating,
		IgnoreRelationshipsWhenMigrating:         conf.IgnoreRelationshipsWhenMigrating,
		DisableNestedTransaction:                 conf.DisableNestedTransaction,
		AllowGlobalUpdate:                        conf.AllowGlobalUpdate,
		QueryFields:                              conf.QueryFields,
		CreateBatchSize:                          conf.CreateBatchSize,
		TranslateError:                           conf.TranslateError,
		Logger:                                   l,
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
