package orm

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/opentelemetry/tracing"
)

// NewBasicMySQL 创建一个简单的mysql连接
func NewBasicMySQL(host, user, pwd, name string) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=%t&loc=%s",
		user,
		pwd,
		host,
		name,
		true,
		"Local")

	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Panicf("open mysql failed. database name: %s, err: %+v", name, err)
	}

	db.Set("gorm:table_options", "CHARSET=utf8mb4")

	return db
}

// NewMySQL 链接数据库，生成数据库实例
func NewMySQL(c *Config) (db *gorm.DB) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=%t&loc=%s",
		c.UserName,
		c.Password,
		c.Addr,
		c.Name,
		true,
		"Local")

	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Panicf("open mysql failed. database name: %s, err: %+v", c.Name, err)
	}
	// set for db connection
	// 用于设置最大打开的连接数，默认值为0表示不限制.设置最大的连接数，可以避免并发太高导致连接mysql出现too many connections的错误。
	sqlDB.SetMaxOpenConns(c.MaxOpenConn)
	// 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
	sqlDB.SetMaxIdleConns(c.MaxIdleConn)
	// 单个连接最大存活时间，建议设置比数据库超时时长(wait_timeout)稍小一些
	sqlDB.SetConnMaxLifetime(c.ConnMaxLifeTime)

	db, err = gorm.Open(mysql.New(mysql.Config{Conn: sqlDB}), gormConfig(c))
	if err != nil {
		log.Panicf("database connection failed. database name: %s, err: %+v", c.Name, err)
	}
	db.Set("gorm:table_options", "CHARSET=utf8mb4")

	if c.Trace { //链路追踪
		if err = db.Use(tracing.NewPlugin(tracing.WithoutMetrics())); err != nil {
			log.Panicf("use tracing failed. database name: %s, err: %+v", c.Name, err)
		}
	}
	return db
}

// gormConfig 根据配置决定是否开启日志
func gormConfig(c *Config) *gorm.Config {
	conf := &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, //禁用自动创建数据库外键约束
		PrepareStmt:                              true, //PreparedStmt 在执行任何 SQL 时都会创建一个 prepared statement 并将其缓存，以提高后续的效率
		Logger:                                   logger.Default.LogMode(logger.Info),
	}
	if c.TablePrefix != "" {
		conf.NamingStrategy = schema.NamingStrategy{
			TablePrefix:   c.TablePrefix, // 表名前缀，`User` 的表名应该是 `t_users`
			SingularTable: true,          // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		}
	}
	// 打印所有SQL
	if c.Debug {
		conf.Logger = logger.Default.LogMode(logger.Info)
	}
	// 只打印慢查询
	if c.SlowThreshold > 0 {
		conf.Logger = logger.New(
			//将标准输出作为Writer
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				//设定慢查询时间阈值
				SlowThreshold: c.SlowThreshold, // nolint: golint
				//设置日志级别，只有指定级别以上会输出慢查询日志
				LogLevel: logger.Warn,
			},
		)
	}
	return conf
}
