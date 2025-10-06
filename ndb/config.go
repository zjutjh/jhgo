package ndb

import (
	"time"

	"gorm.io/gorm/logger"
)

var DefaultConfig = Config{
	Host:      "localhost",
	Port:      3306,
	Database:  "",
	Username:  "",
	Password:  "",
	Charset:   "utf8mb4",
	ParseTime: "True",
	Loc:       "Local",

	DriverName:                    "",
	ServerVersion:                 "",
	SkipInitializeWithVersion:     false,
	DefaultStringSize:             0,
	DefaultDatetimePrecision:      0,
	DisableWithReturning:          false,
	DisableDatetimePrecision:      false,
	DontSupportRenameIndex:        false,
	DontSupportRenameColumn:       false,
	DontSupportForShareClause:     false,
	DontSupportNullAsDefaultValue: false,
	DontSupportRenameColumnUnique: false,

	SkipDefaultTransaction:                   false,
	FullSaveAssociations:                     false,
	DryRun:                                   false,
	PrepareStmt:                              false,
	DisableAutomaticPing:                     false,
	DisableForeignKeyConstraintWhenMigrating: false,
	IgnoreRelationshipsWhenMigrating:         false,
	DisableNestedTransaction:                 false,
	AllowGlobalUpdate:                        false,
	QueryFields:                              false,
	CreateBatchSize:                          0,
	TranslateError:                           true,

	OpenLogger:                true,
	Log:                       "",
	SlowThreshold:             200 * time.Millisecond,
	Colorful:                  false,
	IgnoreRecordNotFoundError: true,
	ParameterizedQueries:      false,
	LogLevel:                  logger.Warn,

	MaxIdleConns:    100,
	MaxOpenConns:    200,
	ConnMaxLifetime: 5 * time.Minute,
	ConnMaxIdleTime: 1 * time.Minute,
}

type Config struct {
	// 基础系列
	Host      string `mapstructure:"host"`
	Port      int    `mapstructure:"port"`
	Database  string `mapstructure:"database"`
	Username  string `mapstructure:"username"`
	Password  string `mapstructure:"password"`
	Charset   string `mapstructure:"charset"`
	ParseTime string `mapstructure:"parse_time"`
	Loc       string `mapstructure:"loc"`

	// gorm MySQL系列
	DriverName                    string `mapstructure:"driver_name"`
	ServerVersion                 string `mapstructure:"server_version"`
	SkipInitializeWithVersion     bool   `mapstructure:"skip_initialize_with_version"`
	DefaultStringSize             uint   `mapstructure:"default_string_size"`
	DefaultDatetimePrecision      int    `mapstructure:"default_datetime_precision"`
	DisableWithReturning          bool   `mapstructure:"disable_with_returning"`
	DisableDatetimePrecision      bool   `mapstructure:"disable_datetime_precision"`
	DontSupportRenameIndex        bool   `mapstructure:"dont_support_rename_index"`
	DontSupportRenameColumn       bool   `mapstructure:"dont_support_rename_column"`
	DontSupportForShareClause     bool   `mapstructure:"dont_support_for_share_clause"`
	DontSupportNullAsDefaultValue bool   `mapstructure:"dont_support_null_as_default_value"`
	DontSupportRenameColumnUnique bool   `mapstructure:"dont_support_rename_column_unique"`

	// gorm系列
	SkipDefaultTransaction                   bool `mapstructure:"skip_default_transaction"`
	FullSaveAssociations                     bool `mapstructure:"full_save_associations"`
	DryRun                                   bool `mapstructure:"dry_run"`
	PrepareStmt                              bool `mapstructure:"prepare_stmt"`
	DisableAutomaticPing                     bool `mapstructure:"disable_automatic_ping"`
	DisableForeignKeyConstraintWhenMigrating bool `mapstructure:"disable_foreign_key_constraint_when_migrating"`
	IgnoreRelationshipsWhenMigrating         bool `mapstructure:"ignore_relationships_when_migrating"`
	DisableNestedTransaction                 bool `mapstructure:"disable_nested_transaction"`
	AllowGlobalUpdate                        bool `mapstructure:"allow_global_update"`
	QueryFields                              bool `mapstructure:"query_fields"`
	CreateBatchSize                          int  `mapstructure:"create_batch_size"`
	TranslateError                           bool `mapstructure:"translate_error"`

	// gorm logger系列
	OpenLogger                bool            `mapstructure:"open_logger"`
	Log                       string          `mapstructure:"log"`
	SlowThreshold             time.Duration   `mapstructure:"slow_threshold"`
	Colorful                  bool            `mapstructure:"colorful"`
	IgnoreRecordNotFoundError bool            `mapstructure:"ignore_record_not_found_error"`
	ParameterizedQueries      bool            `mapstructure:"parameterized_queries"`
	LogLevel                  logger.LogLevel `mapstructure:"log_level"`

	// sql系列
	MaxIdleConns    int           `mapstructure:"max_idle_conns"`
	MaxOpenConns    int           `mapstructure:"max_open_conns"`
	ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime"`
	ConnMaxIdleTime time.Duration `mapstructure:"conn_max_idle_time"`
}
