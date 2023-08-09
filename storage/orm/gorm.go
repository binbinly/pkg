package orm

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/opentelemetry/tracing"
)

const (
	DriverPostgres = "postgres"
	DriverMysql    = "mysql"
)

// NewDB create a db
func NewDB(c *Config) (db *gorm.DB) {
	var sqlDB *sql.DB
	if c.Driver == DriverPostgres {
		sqlDB = openPostgres(c)
	} else {
		sqlDB = openMysql(c)
	}

	// set for db connection
	// 用于设置最大打开的连接数，默认值为0表示不限制.设置最大的连接数，可以避免并发太高导致连接mysql出现too many connections的错误。
	if c.MaxOpenConn > 0 {
		sqlDB.SetMaxOpenConns(c.MaxOpenConn)
	}
	// 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
	if c.MaxIdleConn > 0 {
		sqlDB.SetMaxIdleConns(c.MaxIdleConn)
	}
	// 单个连接最大存活时间，建议设置比数据库超时时长(wait_timeout)稍小一些
	if c.ConnMaxLifeTime > 0 {
		sqlDB.SetConnMaxLifetime(c.ConnMaxLifeTime)
	}

	var err error
	if c.Driver == DriverPostgres {
		db, err = gorm.Open(postgres.New(postgres.Config{Conn: sqlDB, PreferSimpleProtocol: true}), gormConfig(c))
	} else {
		db, err = gorm.Open(mysql.New(mysql.Config{Conn: sqlDB}), gormConfig(c))
	}
	db.Set("gorm:table_options", "CHARSET=utf8mb4")
	if err != nil {
		log.Panicf("database %s connection failed. database name: %s, err: %+v", c.Driver, c.Database, err)
	}

	if c.Trace { //链路追踪
		if err = db.Use(tracing.NewPlugin(tracing.WithoutMetrics())); err != nil {
			log.Panicf("use tracing failed. database name: %s, err: %+v", c.Database, err)
		}
	}

	return db
}

// gormConfig 根据配置决定是否开启日志
func gormConfig(c *Config) *gorm.Config {
	conf := &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, //禁用自动创建数据库外键约束
		PrepareStmt:                              true, //PreparedStmt 在执行任何 SQL 时都会创建一个 prepared statement 并将其缓存，以提高后续的效率
		Logger:                                   logger.Default.LogMode(logger.Warn),
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

func openMysql(c *Config) *sql.DB {
	dsn := c.Dsn
	if dsn == "" {
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=%t&loc=%s",
			c.User, c.Password, c.Host, c.Port, c.Database, true, "Local")
	}
	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Panicf("open %s failed. database name: %s, err: %+v", c.Driver, c.Database, err)
	}

	return sqlDB
}

func openPostgres(c *Config) *sql.DB {
	dsn := c.Dsn
	if dsn == "" {
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
			c.Host, c.User, c.Password, c.Database, c.Port, "disable", "Asia/Shanghai")
	}

	sqlDB, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Panicf("open %s failed. database name: %s, err: %+v", c.Driver, c.Database, err)
	}

	return sqlDB
}
