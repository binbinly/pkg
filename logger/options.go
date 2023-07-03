package logger

var _defOptions = Options{
	JsonEncoding:    true,
	CallerSkipCount: 2,
	Level:           DebugLevel,
	Rotation: Rotation{
		MaxSize:    256,
		MaxBackups: 300,
		MaxAge:     30,
		LocalTime:  true,
		Compress:   true,
	},
}

type Option func(*Options)

type Options struct {
	Debug           bool
	DisableConsole  bool
	JsonEncoding    bool
	CallerSkipCount int
	LogDir          string
	Level           string
	Fields          map[string]any
	Rotation        Rotation
}

type Rotation struct {
	MaxSize    int  // 单个文件最大尺寸，默认单位 M
	MaxBackups int  // 最多保留 300 个备份
	MaxAge     int  // 最大时间，默认单位 day
	LocalTime  bool // 使用本地时间
	Compress   bool // 是否压缩 disabled by default
}

func WithDebug() Option {
	return func(args *Options) {
		args.Debug = true
	}
}

func WithDisableConsole() Option {
	return func(args *Options) {
		args.DisableConsole = true
	}
}

func WithLogDir(dir string) Option {
	return func(args *Options) {
		args.LogDir = dir
	}
}

func WithJsonEncoding() Option {
	return func(args *Options) {
		args.JsonEncoding = true
	}
}

func WithFields(fields map[string]any) Option {
	return func(args *Options) {
		args.Fields = fields
	}
}

func WithLevel(level string) Option {
	return func(args *Options) {
		args.Level = level
	}
}

func WithCallerSkipCount(c int) Option {
	return func(args *Options) {
		args.CallerSkipCount = c
	}
}

func WithRotation(r Rotation) Option {
	return func(args *Options) {
		args.Rotation = r
	}
}
