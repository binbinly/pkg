package orm

import "time"

// Config mysql config
type Config struct {
	Name            string
	Addr            string
	UserName        string
	Password        string
	TablePrefix     string
	Debug           bool
	Trace           bool
	MaxIdleConn     int
	MaxOpenConn     int
	ConnMaxLifeTime time.Duration
	SlowThreshold   time.Duration // 慢查询时长，默认500ms
}
