/*
 * @Auther: Biny
 * @Description:
 * @Date: 2022-01-23 02:18:58
 * @LastEditTime: 2022-01-25 10:13:44
 */
// /*
//  * @Auther: Biny
//  * @Description:
//  * @Date: 2022-01-23 02:18:58
//  * @LastEditTime: 2022-01-23 03:07:42
//  */
package models

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

//var sqlDB *sql.DB

func init() {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,       // 禁用彩色打印
		},
	)
	//DSN 全称叫 Data Source Name，数据库的源名称
	dsn := fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=%t&loc=%s&timeout=3s`, "root", "root", "192.168.31.22", "tp6-admin", true, "Local")
	//dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", "root", "root", "192.168.31.22", 3306, "tp6-admin")
	dbNew, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})

	if err != nil {
		panic("数据库连接失败")
	}
	db = dbNew

	//db.Callback().Query().After("gorm:query").Register("my_plugin:after_query", afterQuery)

	//???连接池???
	// sqlDbNew, err := db.DB()
	// if err != nil {
	// 	panic(err.Error())
	// }

	// // SetMaxOpenConns sets the maximum number of open connections to the database.
	// sqlDbNew.SetMaxOpenConns(100)

	// // SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	// sqlDbNew.SetConnMaxLifetime(time.Duration(10) * time.Second)

	// // SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	// sqlDbNew.SetMaxIdleConns(100)

	// sqlDB = sqlDbNew
}

//, NowFunc: timeformat
// func timeformat() time.Time {
// 	fmt.Println("-------timeformat--------")
// 	return time.Now().Local().Truncate(time.Microsecond)
// }

// func afterQuery(db *gorm.DB) {
// 	fmt.Println("-------afterQuery-----")
// 	if db.Error == nil && db.Statement.Schema != nil && !db.Statement.SkipHooks && db.Statement.Schema.AfterFind && db.RowsAffected > 0 {
// 		fmt.Println("-------afterQuery1-----", db)
// 		// fmt.Println("-------afterQuery2-----",db)
// 		// callMethod(db, func(value interface{}, tx *gorm.DB) bool {
// 		// 	if i, ok := value.(AfterFindInterface); ok {
// 		// 		db.AddError(i.AfterFind(tx))
// 		// 		return true
// 		// 	}
// 		// 	return false
// 		// })
// 	}
// }

// type AfterFindInterface interface {
// 	AfterFind(*gorm.DB) error
// }
// import (
// 	"fmt"
// 	"sync"
// 	"time"

// 	"github.com/mygin/pkg/db"
// 	"gorm.io/gorm"
// )

// var (
// 	once sync.Once
// )

// type MySQLOptions struct {
// 	Host                  string        `json:"host,omitempty"                     mapstructure:"host"`
// 	Username              string        `json:"username,omitempty"                 mapstructure:"username"`
// 	Password              string        `json:"-"                                  mapstructure:"password"`
// 	Database              string        `json:"database"                           mapstructure:"database"`
// 	MaxIdleConnections    int           `json:"max-idle-connections,omitempty"     mapstructure:"max-idle-connections"`
// 	MaxOpenConnections    int           `json:"max-open-connections,omitempty"     mapstructure:"max-open-connections"`
// 	MaxConnectionLifeTime time.Duration `json:"max-connection-life-time,omitempty" mapstructure:"max-connection-life-time"`
// 	LogLevel              int           `json:"log-level"                          mapstructure:"log-level"`
// }

// // NewMySQLOptions create a `zero` value instance.
// func NewMySQLOptions() *MySQLOptions {
// 	return &MySQLOptions{
// 		Host:                  "192.168.31.22:3306",
// 		Username:              "root",
// 		Password:              "root",
// 		Database:              "tp6-admin",
// 		MaxIdleConnections:    100,
// 		MaxOpenConnections:    100,
// 		MaxConnectionLifeTime: time.Duration(10) * time.Second,
// 	}
// }

// // GetMySQLFactoryOr create mysql factory with the given config.
// func GetMySQLFactoryOr(opts *MySQLOptions) (*gorm.DB, error) {
// 	if opts == nil {
// 		return nil, fmt.Errorf("failed to get mysql Options")
// 	}

// 	var err error
// 	var dbIns *gorm.DB
// 	once.Do(func() {
// 		options := &db.Options{
// 			Host:                  opts.Host,
// 			Username:              opts.Username,
// 			Password:              opts.Password,
// 			Database:              opts.Database,
// 			MaxIdleConnections:    opts.MaxIdleConnections,
// 			MaxOpenConnections:    opts.MaxOpenConnections,
// 			MaxConnectionLifeTime: opts.MaxConnectionLifeTime,
// 		}
// 		dbIns, err = db.New(options)

// 		// uncomment the following line if you need auto migration the given models
// 		// not suggested in production environment.
// 		// migrateDatabase(dbIns)

// 	})

// 	if dbIns == nil || err != nil {
// 		return nil, fmt.Errorf("failed to get mysql : %+v, error: %w", dbIns, err)
// 	}

// 	return dbIns, nil
// }

type MyTime time.Time

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
