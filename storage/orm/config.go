package orm

import "time"

// Config mysql config
type Config struct {
	Addr            string
	User            string
	Password        string
	Database        string
	TablePrefix     string
	Debug           bool
	Trace           bool
	MaxIdleConn     int
	MaxOpenConn     int
	ConnMaxLifeTime time.Duration
	SlowThreshold   time.Duration // 慢查询时长，默认500ms
}
