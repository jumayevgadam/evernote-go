package logger

import (
	"os"

	"github.com/jumayevgadam/evernote-go/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger methods interface.
type Logger interface {
	InitLogger()
	Debug(args ...interface{})
	Debugf(template string, args ...interface{})
	Info(args ...interface{})
	Infof(template string, args ...interface{})
	Warn(args ...interface{})
	Warnf(template string, args ...interface{})
	Error(args ...interface{})
	Errorf(template string, args ...interface{})
	DPanic(args ...interface{})
	DPanicf(template string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(template string, args ...interface{})
}

// APILogger struct.
type APILogger struct {
	cfg         *config.Config
	sugarLogger *zap.SugaredLogger
}

// NewAPILogger func creates new logger.
func NewAPILogger(cfg *config.Config) *APILogger {
	return &APILogger{
		cfg: cfg,
	}
}

var loggerLevelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

func (l *APILogger) getLoggerLevel(cfg *config.Config) zapcore.Level {
	level, exist := loggerLevelMap[cfg.Logger.Level]
	if !exist {
		return zapcore.DebugLevel
	}

	return level
}

// InitLogger func initializes logger.
func (l *APILogger) InitLogger() {
	logLevel := l.getLoggerLevel(l.cfg)

	logWriter := zapcore.AddSync(os.Stderr)

	var encoderCfg zapcore.EncoderConfig
	if l.cfg.Logger.Development {
		encoderCfg = zap.NewDevelopmentEncoderConfig()
	} else {
		encoderCfg = zap.NewProductionEncoderConfig()
	}

	var encoder zapcore.Encoder
	encoderCfg.LevelKey = "LEVEL"
	encoderCfg.CallerKey = "CALLER"
	encoderCfg.TimeKey = "TIME"
	encoderCfg.MessageKey = "MESSAGE"
	encoderCfg.NameKey = "NAME"

	if l.cfg.Logger.Encoding == "json" {
		encoder = zapcore.NewJSONEncoder(encoderCfg)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderCfg)
	}

	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(encoder, logWriter, zap.NewAtomicLevelAt(logLevel))
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	l.sugarLogger = logger.Sugar()
	if err := l.sugarLogger.Sync(); err != nil {
		l.sugarLogger.Error(err)
	}
}

// Debug func logs debug message.
func (l *APILogger) Debug(args ...interface{}) {
	l.sugarLogger.Debug(args...)
}

// Debugf func logs debug message with format.
func (l *APILogger) Debugf(template string, args ...interface{}) {
	l.sugarLogger.Debugf(template, args...)
}

// Info func logs info message.
func (l *APILogger) Info(args ...interface{}) {
	l.sugarLogger.Info(args...)
}

// Infof func logs info message with format.
func (l *APILogger) Infof(template string, args ...interface{}) {
	l.sugarLogger.Infof(template, args...)
}

// Warn func logs warn message.
func (l *APILogger) Warn(args ...interface{}) {
	l.sugarLogger.Warn(args...)
}

// Warnf func logs warn message with format.
func (l *APILogger) Warnf(template string, args ...interface{}) {
	l.sugarLogger.Warnf(template, args...)
}

// Error func logs error message.
func (l *APILogger) Error(args ...interface{}) {
	l.sugarLogger.Error(args...)
}

// Errorf func logs error message with format.
func (l *APILogger) Errorf(template string, args ...interface{}) {
	l.sugarLogger.Errorf(template, args...)
}

// DPanic func logs panic message.
func (l *APILogger) DPanic(args ...interface{}) {
	l.sugarLogger.DPanic(args...)
}

// DPanicf func logs panic message with format.
func (l *APILogger) DPanicf(template string, args ...interface{}) {
	l.sugarLogger.DPanicf(template, args...)
}

// Panic func logs panic message.
func (l *APILogger) Panic(args ...interface{}) {
	l.sugarLogger.Panic(args...)
}

// Panicf func logs panic message with format.
func (l *APILogger) Panicf(template string, args ...interface{}) {
	l.sugarLogger.Panicf(template, args...)
}

// Fatal func logs fatal message.
func (l *APILogger) Fatal(args ...interface{}) {
	l.sugarLogger.Fatal(args...)
}

// Fatalf func logs fatal message with format.
func (l *APILogger) Fatalf(template string, args ...interface{}) {
	l.sugarLogger.Fatalf(template, args...)
}
