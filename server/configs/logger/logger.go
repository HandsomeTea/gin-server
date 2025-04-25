package logger

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"gin-server/server/configs/env"
	"gin-server/server/globals"
)

var logLevels = map[string]zapcore.Level{
	"debug":  zap.DebugLevel,
	"info":   zap.InfoLevel,
	"warn":   zap.WarnLevel,
	"error":  zap.ErrorLevel,
	"dpanic": zap.DPanicLevel,
	"panic":  zap.PanicLevel,
	"fatal":  zap.FatalLevel,
}

type loggerConfig struct {
	Name         string
	DefaultLevel zapcore.Level
	Level        string
	ShowCaller   bool // 是否显示调用位置
	CustomFormat func(zapcore.Entry, string) string
}

const timeFormat = "2006-01-02 15:04:05.000 Z07:00"

func createLogger(config loggerConfig) *zap.Logger {
	encoderConfig := zapcore.EncoderConfig{}
	if config.ShowCaller {
		encoderConfig.CallerKey = "caller"
	}

	level := zap.NewAtomicLevel()
	if parsedLevel, exists := logLevels[config.Level]; exists {
		level.SetLevel(parsedLevel)
	} else {
		level.SetLevel(config.DefaultLevel) // 默认级别回退
	}

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.WriteSyncer(os.Stdout),
		level,
	)

	opts := []zap.Option{zap.Hooks(func(entry zapcore.Entry) error {
		fmt.Print(config.CustomFormat(entry, config.Name))
		return nil
	})}

	if config.ShowCaller {
		opts = append(opts, zap.AddCaller())
	}

	return zap.New(core, opts...)
}

func customDevLogFormat(entry zapcore.Entry, logName string) string {
	timeStr := fmt.Sprintf("[%s]", entry.Time.Format(timeFormat))
	levelStr := fmt.Sprintf("[%s]", strings.ToUpper(entry.Level.String()))
	nameStr := fmt.Sprintf("[%s:%s]", globals.SERVER_NAME, strings.ToUpper(logName))
	callerStr := fmt.Sprintf("[%s:%d]", entry.Caller.File, entry.Caller.Line)
	msgStr := entry.Message + "\n"

	var line string
	switch entry.Level {
	case zapcore.DebugLevel:
		line = color.CyanString("%s %s %s %s %s", timeStr, levelStr, nameStr, callerStr, msgStr)
	case zapcore.InfoLevel:
		line = color.GreenString("%s %s %s %s %s", timeStr, levelStr, nameStr, callerStr, msgStr)
	case zapcore.WarnLevel:
		line = color.YellowString("%s %s %s %s %s", timeStr, levelStr, nameStr, callerStr, msgStr)
	case zapcore.ErrorLevel:
		line = color.RedString("%s %s %s %s %s", timeStr, levelStr, nameStr, callerStr, msgStr)
	case zapcore.FatalLevel:
		line = color.MagentaString("%s %s %s %s %s", timeStr, levelStr, nameStr, callerStr, msgStr)
	case zapcore.PanicLevel:
		line = color.HiMagentaString("%s %s %s %s %s", timeStr, levelStr, nameStr, callerStr, msgStr)
	default:
		line = fmt.Sprintf("%s %s %s %s %s", timeStr, levelStr, nameStr, callerStr, msgStr)
	}

	return line + "\n"
}

var Log *zap.Logger

func customTraceLogFormat(entry zapcore.Entry, logName string) string {
	timeStr := fmt.Sprintf("[%s]", entry.Time.Format(timeFormat))
	levelStr := fmt.Sprintf("[%s]", strings.ToUpper(entry.Level.String()))
	nameStr := fmt.Sprintf("[%s:%s]", globals.SERVER_NAME, strings.ToUpper(logName))
	msgStr := entry.Message + "\n"

	var line string
	switch entry.Level {
	case zapcore.DebugLevel:
		line = color.CyanString("%s %s %s %s", timeStr, levelStr, nameStr, msgStr)
	case zapcore.InfoLevel:
		line = color.GreenString("%s %s %s %s", timeStr, levelStr, nameStr, msgStr)
	case zapcore.WarnLevel:
		line = color.YellowString("%s %s %s %s", timeStr, levelStr, nameStr, msgStr)
	case zapcore.ErrorLevel:
		line = color.RedString("%s %s %s %s", timeStr, levelStr, nameStr, msgStr)
	case zapcore.FatalLevel:
		line = color.MagentaString("%s %s %s %s", timeStr, levelStr, nameStr, msgStr)
	case zapcore.PanicLevel:
		line = color.HiMagentaString("%s %s %s %s", timeStr, levelStr, nameStr, msgStr)
	default:
		line = fmt.Sprintf("%s %s %s %s", timeStr, levelStr, nameStr, msgStr)
	}

	return line + "\n"
}

var TraceLog *zap.Logger

func customSystemLoggerFormat(entry zapcore.Entry, logName string) string {
	timeStr := fmt.Sprintf("[%s]", entry.Time.Format(timeFormat))
	levelStr := fmt.Sprintf("[%s]", strings.ToUpper(entry.Level.String()))
	nameStr := fmt.Sprintf("[%s:%s]", globals.SERVER_NAME, strings.ToUpper(logName))
	msgStr := entry.Message + "\n"

	var line string
	switch entry.Level {
	case zapcore.DebugLevel:
		line = color.CyanString("%s %s %s %s", timeStr, levelStr, nameStr, msgStr)
	case zapcore.InfoLevel:
		line = color.GreenString("%s %s %s %s", timeStr, levelStr, nameStr, msgStr)
	case zapcore.WarnLevel:
		line = color.YellowString("%s %s %s %s", timeStr, levelStr, nameStr, msgStr)
	case zapcore.ErrorLevel:
		line = color.RedString("%s %s %s %s", timeStr, levelStr, nameStr, msgStr)
	case zapcore.FatalLevel:
		line = color.MagentaString("%s %s %s %s", timeStr, levelStr, nameStr, msgStr)
	case zapcore.PanicLevel:
		line = color.HiMagentaString("%s %s %s %s", timeStr, levelStr, nameStr, msgStr)
	default:
		line = fmt.Sprintf("%s %s %s %s", timeStr, levelStr, nameStr, msgStr)
	}

	return line + "\n"
}

var SystemLog *zap.Logger

func Close() {
	SystemLog.Sync()
	Log.Sync()
	TraceLog.Sync()
}

func init() {
	logLevel, _ := env.GetEnv("LOG_LEVEL")
	traceLogLevel, _ := env.GetEnv("TRACE_LOG_LEVEL")

	SystemLog = createLogger(loggerConfig{
		Name:         "system",
		DefaultLevel: zapcore.DebugLevel,
		CustomFormat: customSystemLoggerFormat,
	})
	Log = createLogger(loggerConfig{
		Name:         "dev",
		DefaultLevel: zapcore.DebugLevel,
		Level:        logLevel,
		ShowCaller:   true,
		CustomFormat: customDevLogFormat,
	})
	TraceLog = createLogger(loggerConfig{
		Name:         "trace",
		DefaultLevel: zapcore.DebugLevel,
		Level:        traceLogLevel,
		CustomFormat: customTraceLogFormat,
	})
}
