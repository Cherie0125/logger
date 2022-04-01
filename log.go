package logger

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

// log object
var logger *Logger

// Encapsulating log objects
type Logger struct {
	coreLogger *zap.SugaredLogger
}

// function for init logger
func InitLogger(config *LoggerConfig) {
	logger = NewLogger(config)
}

// function for create logger object
func NewLogger(config *LoggerConfig) *Logger {
	newLogger := new(Logger)

	var allCore []zapcore.Core
	var level zapcore.Level

	// make true log level
	switch config.LogLevel {
	case ErrorLevel:
		level = zap.ErrorLevel
	case InfoLevel:
		level = zap.InfoLevel
	case DebugLevel:
		level = zap.ErrorLevel
	}

	// get encoder
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	// add file log output
	if config.File == true {
		// get writeSyncer
		lumberJackLogger := &lumberjack.Logger{
			Filename:   config.Fileconfig.Filename,
			MaxSize:    config.Fileconfig.MaxSize,
			MaxAge:     config.Fileconfig.MaxAge,
			MaxBackups: config.Fileconfig.MaxBackups,
			LocalTime:  config.Fileconfig.LocalTime,
			Compress:   config.Fileconfig.Compress,
		}
		fileWriter := zapcore.AddSync(lumberJackLogger)
		allCore = append(allCore, zapcore.NewCore(encoder, fileWriter, level))
	}

	// add console log output
	if config.Console == true {
		consoleWriter := zapcore.Lock(os.Stdout)
		allCore = append(allCore, zapcore.NewCore(encoder, consoleWriter, level))
	}

	core := zapcore.NewTee(allCore...)
	oriLogger := zap.New(core, zap.AddCaller())
	newLogger.coreLogger = oriLogger.Sugar()

	return newLogger
}

func Error(args ...interface{}) {
	logger.coreLogger.Error(args...)
}

func Errorf(pattern string, args ...interface{}) {
	logger.coreLogger.Errorf(pattern, args)
}

func Info(args ...interface{}) {
	logger.coreLogger.Info(args...)
}

func Infof(pattern string, args ...interface{}) {
	logger.coreLogger.Infof(pattern, args...)
}

func Debug(args ...interface{}) {
	logger.coreLogger.Debug(args...)
}

func Debugf(pattern string, args ...interface{}) {
	logger.coreLogger.Debugf(pattern, args...)
}
