package logger

type Level int8

// log level
const (
	ErrorLevel Level = iota - 1
	InfoLevel
	DebugLevel

	_minLevel = ErrorLevel
	_maxLevel = DebugLevel
)

// the configure of writing logs to files
type FileConfig struct {
	Filename string // log file's name and path
	MaxSize int
	MaxAge int
	MaxBackups int
	LocalTime bool
	Compress bool
}

// the configure of log object
type LoggerConfig struct {
	LogLevel Level
	Color bool
	Console bool
	File bool
	Fileconfig *FileConfig
}

// function for create log file configure
func NewLoggerFileConfig() *FileConfig {
	fileConfig := new(FileConfig)

	return fileConfig
}

// function for create log configure
func NewLoggerConfilg() *LoggerConfig {
	config := new(LoggerConfig)

	return config
}