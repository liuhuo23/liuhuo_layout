package data

import (
	"context"
	"fmt"
	"liuhuo23/liuos/internal/conf"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	grom_log "gorm.io/gorm/logger"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewUserData)

// Data .
type loggerAdapter struct {
	logger log.Logger
}

func (l *loggerAdapter) LogMode(level grom_log.LogLevel) grom_log.Interface {
	return l
}

func (l *loggerAdapter) Info(ctx context.Context, msg string, data ...interface{}) {
	log.NewHelper(l.logger).Infof(msg, data...)
}

func (l *loggerAdapter) Warn(ctx context.Context, msg string, data ...interface{}) {
	log.NewHelper(l.logger).Warnf(msg, data...)
}

func (l *loggerAdapter) Error(ctx context.Context, msg string, data ...interface{}) {
	log.NewHelper(l.logger).Errorf(msg, data...)
}

func (l *loggerAdapter) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, rows := fc()
	elapsed := time.Since(begin)
	if err != nil {
		log.NewHelper(l.logger).Errorf("SQL: %s, Error: %v, Elapsed: %v, Rows: %d", sql, err, elapsed, rows)
	} else {
		log.NewHelper(l.logger).Infof("SQL: %s, Elapsed: %v, Rows: %d", sql, elapsed, rows)
	}
}

type Data struct {
	// TODO wrapped database client
	Db *gorm.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	fmt.Println("hello")
	log.NewHelper(logger).Infof("connecting to database: %s", c.Database.Source)

	db, err := gorm.Open(mysql.Open(c.Database.Source), &gorm.Config{
		Logger: &loggerAdapter{logger: logger},
	})
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %v", err))
	}
	cleanup := func() {

		sqlDB, err := db.DB()
		if err != nil {
			log.NewHelper(logger).Error(err)
		} else {
			if err := sqlDB.Close(); err != nil {
				log.NewHelper(logger).Error(err)
			}
		}
	}
	return &Data{Db: db}, cleanup, nil
}
