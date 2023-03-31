package model

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/mwqnice/oh-admin/global"
	"github.com/mwqnice/oh-admin/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"strings"
	"time"
)

const (
	STATE_OPEN  = 1
	STATE_CLOSE = 0
)

//MyTime 自定义时间
type MyTime time.Time

type Model struct {
	ID          int    `gorm:"primary_key" json:"id"`
	CreatedTime MyTime `json:"created_time" gorm:"column:created_time"`
	UpdatedTime MyTime `json:"updated_time" gorm:"column:updated_time"`
}

func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	s := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
	)

	loggerMode := logger.Info
	if global.ServerSetting.RunMode == "debug" {
		loggerMode = logger.Info
	}
	lw := LoggerWriter{}
	logs := logger.New(lw.New(loggerMode), logger.Config{LogLevel: loggerMode})
	db, err := gorm.Open(mysql.Open(s), &gorm.Config{
		Logger: logs,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		return nil, err
	}

	dB, err := db.DB()
	if err != nil {
		return nil, err
	}

	dB.SetMaxOpenConns(databaseSetting.MaxOpenConns)
	dB.SetMaxIdleConns(databaseSetting.MaxIdleConns)

	return db, nil
}

func (m *Model) BeforeUpdate(db *gorm.DB) error {
	n := &MyTime{}
	if err := n.Scan(time.Now()); err != nil {
		return err
	}
	db.Statement.SetColumn("updated_time", *n)

	return nil
}

func (m *Model) BeforeCreate(db *gorm.DB) error {
	n := &MyTime{}
	if err := n.Scan(time.Now()); err != nil {
		return err
	}
	db.Statement.SetColumn("created_time", *n)
	db.Statement.SetColumn("updated_time", *n)

	return nil
}
func (t *MyTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var err error
	//前端接收的时间字符串
	str := string(data)
	//去除接收的str收尾多余的"
	timeStr := strings.Trim(str, "\"")
	t1, err := time.Parse("2006-01-02 15:04:05", timeStr)
	*t = MyTime(t1)
	return err
}

func (t MyTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%v\"", time.Time(t).Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}

func (t MyTime) Value() (driver.Value, error) {
	// MyTime 转换成 time.Time 类型
	tTime := time.Time(t)
	return tTime.Format("2006-01-02 15:04:05"), nil
}

func (t *MyTime) Scan(v interface{}) error {
	switch vt := v.(type) {
	case time.Time:
		// 字符串转成 time.Time 类型
		*t = MyTime(vt)
	default:
		return errors.New("类型处理错误")
	}
	return nil
}

func (t *MyTime) String() string {
	return fmt.Sprintf("hhh:%s", time.Time(*t).String())
}

func (t *MyTime) MyTime2Date() string {
	return time.Time(*t).Format("2006-01-02")
}

type LoggerWriter struct {
	ctx      context.Context
	logLevel logger.LogLevel
}

func (lw LoggerWriter) Printf(msg string, v ...interface{}) {
	global.Logger.Channel("sql_print").Infof(context.Background(), msg, v...)

	if lw.logLevel == logger.Info {
		soLog := log.New(os.Stdout, "\r\n", log.LstdFlags)
		soLog.Printf(msg, v...)
	}
}

func (lw *LoggerWriter) New(loggerLevel logger.LogLevel) LoggerWriter {
	lw.logLevel = loggerLevel
	return *lw
}
