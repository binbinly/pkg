package logger

import (
	"fmt"
	"io"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type zapLogger struct {
	zap  *zap.Logger
	opts Options
}

func (l *zapLogger) Init(opts ...Option) error {
	for _, o := range opts {
		o(&l.opts)
	}

	stdout := zapcore.Lock(os.Stdout) // lock for concurrent safe
	stderr := zapcore.Lock(os.Stderr) // lock for concurrent safe

	infoEnabler := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= getZapLevel(l.opts.Level) && lvl < zapcore.WarnLevel
	})
	warnEnabler := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.WarnLevel
	})

	encoderCfg := zap.NewProductionEncoderConfig()
	if l.opts.Debug {
		encoderCfg = zap.NewDevelopmentEncoderConfig()
	}

	encoder := zapcore.NewConsoleEncoder(encoderCfg)
	if l.opts.JsonEncoding {
		encoder = zapcore.NewJSONEncoder(encoderCfg)
	}

	var cores []zapcore.Core
	if !l.opts.DisableConsole {
		consoleCore := zapcore.NewCore(
			encoder,
			zapcore.NewMultiWriteSyncer(stdout),
			infoEnabler,
		)
		errorCore := zapcore.NewCore(
			encoder,
			zapcore.NewMultiWriteSyncer(stderr),
			warnEnabler,
		)
		cores = append(cores, consoleCore, errorCore)
	}

	if l.opts.LogDir != "" {
		if err := os.MkdirAll(l.opts.LogDir, 0766); err != nil {
			return err
		}
		infoFile := zapcore.AddSync(l.filesizeRotation(l.opts.LogDir + "info.log"))
		warnFile := zapcore.AddSync(l.filesizeRotation(l.opts.LogDir + "warn.log"))
		infoCore := zapcore.NewCore(
			encoder, zapcore.NewMultiWriteSyncer(infoFile),
			infoEnabler,
		)
		warnCore := zapcore.NewCore(
			encoder, zapcore.NewMultiWriteSyncer(warnFile),
			warnEnabler,
		)
		cores = append(cores, infoCore, warnCore)
	}

	// 构造日志
	l.zap = zap.New(
		zapcore.NewTee(cores...),
		zap.AddCaller(),
		zap.AddCallerSkip(l.opts.CallerSkipCount),
		// 堆栈信息
		zap.AddStacktrace(zapcore.ErrorLevel),
		zap.ErrorOutput(stderr),
	)

	if l.opts.Fields != nil {
		fields := make([]zap.Field, 0, len(l.opts.Fields))
		for k, v := range l.opts.Fields {
			fields = append(fields, zap.Any(k, v))
		}
		l.zap = l.zap.With(fields...)
	}
	return nil
}

func (l *zapLogger) Fields(fields map[string]any) Logger {
	data := make([]zap.Field, 0, len(fields))
	for k, v := range fields {
		data = append(data, zap.Any(k, v))
	}

	zl := &zapLogger{
		zap:  l.zap.With(data...),
		opts: l.opts,
	}

	return zl
}

func (l *zapLogger) Log(level string, v ...any) {
	l.zap.Log(getZapLevel(level), fmt.Sprint(v...))
}

func (l *zapLogger) Logf(level string, format string, v ...any) {
	l.zap.Log(getZapLevel(level), fmt.Sprintf(format, v...))
}

func (l *zapLogger) filesizeRotation(file string) io.Writer {
	return &lumberjack.Logger{
		Filename:   file,
		MaxSize:    l.opts.Rotation.MaxSize,
		MaxBackups: l.opts.Rotation.MaxBackups,
		MaxAge:     l.opts.Rotation.MaxAge,
		LocalTime:  l.opts.Rotation.LocalTime,
		Compress:   l.opts.Rotation.Compress,
	}
}

// NewExampleLogger builds a Logger that's designed for use in zap's testable
func NewExampleLogger() Logger {
	return &zapLogger{
		zap: zap.NewExample(),
	}
}

// NewLogger New builds a new logger based on options.
func NewLogger(opts ...Option) Logger {
	l := &zapLogger{opts: _defOptions}
	if err := l.Init(opts...); err != nil {
		return nil
	}

	return l
}

// InitLogger init customize logger
func InitLogger(opts ...Option) {
	l := &zapLogger{opts: _defOptions}
	if err := l.Init(opts...); err != nil {
		panic(err)
	}

	logger = l
}
