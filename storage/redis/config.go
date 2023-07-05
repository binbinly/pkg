package redis

import "time"

// Config redis config
type Config struct {
	Url          string
	Addr         string
	Username     string
	Password     string
	DB           int
	MinIdleConn  int
	PoolSize     int
	DialTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	PoolTimeout  time.Duration
	Trace        bool
}
