package cache

import (
	"time"

	"github.com/binbinly/pkg/codec"
)

type Option func(*Options)

type Options struct {
	expire time.Duration
	codec  codec.Encoding
	prefix string
}

func NewOptions(opt ...Option) Options {
	opts := Options{
		expire: DefaultExpireTime,
		prefix: DefaultPrefix,
		codec:  codec.JSONEncoding{},
	}

	for _, o := range opt {
		o(&opts)
	}

	return opts
}

func WithExpire(d time.Duration) Option {
	return func(o *Options) {
		o.expire = d
	}
}

func WithPrefix(prefix string) Option {
	return func(o *Options) {
		o.prefix = prefix
	}
}

func WithCodec(codec codec.Encoding) Option {
	return func(o *Options) {
		o.codec = codec
	}
}
