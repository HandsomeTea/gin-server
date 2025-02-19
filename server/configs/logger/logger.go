package logger

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	env "gin-server/server/configs/env"
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

func customDevLogFormat(entry zapcore.Entry) string {
	timeStr := fmt.Sprintf("[%s]", entry.Time.Format("2006-01-02 15:04:05.000"))
	levelStr := fmt.Sprintf("[%s]", strings.ToUpper(entry.Level.String())) // 将日志级别转换为大写
	callerStr := fmt.Sprintf("[%s:%d]", entry.Caller.File, entry.Caller.Line)
	msgStr := entry.Message + "\n"

	var line string
	switch entry.Level {
	case zapcore.DebugLevel:
		line = color.CyanString("%s %s %s %s", timeStr, levelStr, callerStr, msgStr)
	case zapcore.InfoLevel:
		line = color.GreenString("%s %s %s %s", timeStr, levelStr, callerStr, msgStr)
	case zapcore.WarnLevel:
		line = color.YellowString("%s %s %s %s", timeStr, levelStr, callerStr, msgStr)
	case zapcore.ErrorLevel:
		line = color.RedString("%s %s %s %s", timeStr, levelStr, callerStr, msgStr)
	case zapcore.FatalLevel:
		line = color.MagentaString("%s %s %s %s", timeStr, levelStr, callerStr, msgStr)
	case zapcore.PanicLevel:
		line = color.HiMagentaString("%s %s %s %s", timeStr, levelStr, callerStr, msgStr)
	default:
		line = fmt.Sprintf("%s %s %s %s", timeStr, levelStr, callerStr, msgStr) // 默认颜色
	}

	return line + "\n" // 添加换行符
}

var Log *zap.Logger

func createDevLogger() {
	encoderConfig := zapcore.EncoderConfig{
		CallerKey: "caller",
	}

	level := zap.NewAtomicLevel()
	logLevel := env.GetEnv("LOG_LEVEL")
	level.SetLevel(logLevels[logLevel])

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.WriteSyncer(os.Stdout),
		level,
	)

	Log = zap.New(core, zap.AddCaller(), zap.Hooks(func(entry zapcore.Entry) error {
		fmt.Print(customDevLogFormat(entry))
		return nil
	}))

	defer Log.Sync()
}

func customTraceLogFormat(entry zapcore.Entry) string {
	timeStr := fmt.Sprintf("[%s]", entry.Time.Format("2006-01-02 15:04:05.000 Z07:00"))
	levelStr := fmt.Sprintf("[%s]", strings.ToUpper(entry.Level.String())) // 将日志级别转换为大写
	msgStr := entry.Message + "\n"

	var line string
	switch entry.Level {
	case zapcore.DebugLevel:
		line = color.CyanString("%s %s %s", timeStr, levelStr, msgStr)
	case zapcore.InfoLevel:
		line = color.GreenString("%s %s %s", timeStr, levelStr, msgStr)
	case zapcore.WarnLevel:
		line = color.YellowString("%s %s %s", timeStr, levelStr, msgStr)
	case zapcore.ErrorLevel:
		line = color.RedString("%s %s %s", timeStr, levelStr, msgStr)
	case zapcore.FatalLevel:
		line = color.MagentaString("%s %s %s", timeStr, levelStr, msgStr)
	case zapcore.PanicLevel:
		line = color.HiMagentaString("%s %s %s", timeStr, levelStr, msgStr)
	default:
		line = fmt.Sprintf("%s %s %s", timeStr, levelStr, msgStr) // 默认颜色
	}

	return line + "\n" // 添加换行符
}

var TraceLog *zap.Logger

func createTraceLogger() {
	encoderConfig := zapcore.EncoderConfig{}

	level := zap.NewAtomicLevel()
	traceLogLevel := env.GetEnv("TRACE_LOG_LEVEL")
	level.SetLevel(logLevels[traceLogLevel])

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.WriteSyncer(os.Stdout),
		level,
	)

	TraceLog = zap.New(core, zap.Hooks(func(entry zapcore.Entry) error {
		fmt.Print(customTraceLogFormat(entry))
		return nil
	}))

	defer TraceLog.Sync()
}

func customSystemLoggerFormat(entry zapcore.Entry) string {
	timeStr := fmt.Sprintf("[%s]", entry.Time.Format("2006-01-02 15:04:05.000"))
	levelStr := fmt.Sprintf("[%s]", strings.ToUpper(entry.Level.String())) // 将日志级别转换为大写
	msgStr := entry.Message + "\n"

	var line string
	switch entry.Level {
	case zapcore.DebugLevel:
		line = color.CyanString("%s %s [SYSTEM] %s", timeStr, levelStr, msgStr)
	case zapcore.InfoLevel:
		line = color.GreenString("%s %s [SYSTEM] %s", timeStr, levelStr, msgStr)
	case zapcore.WarnLevel:
		line = color.YellowString("%s %s [SYSTEM] %s", timeStr, levelStr, msgStr)
	case zapcore.ErrorLevel:
		line = color.RedString("%s %s [SYSTEM] %s", timeStr, levelStr, msgStr)
	case zapcore.FatalLevel:
		line = color.MagentaString("%s %s [SYSTEM] %s", timeStr, levelStr, msgStr)
	case zapcore.PanicLevel:
		line = color.HiMagentaString("%s %s [SYSTEM] %s", timeStr, levelStr, msgStr)
	default:
		line = fmt.Sprintf("%s %s [SYSTEM] %s", timeStr, levelStr, msgStr) // 默认颜色
	}

	return line + "\n" // 添加换行符
}

var SystemLog *zap.Logger

func createSystemLogger() {
	encoderConfig := zapcore.EncoderConfig{}

	level := zap.NewAtomicLevel()
	level.SetLevel(logLevels["debug"])

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.WriteSyncer(os.Stdout),
		level,
	)

	SystemLog = zap.New(core, zap.Hooks(func(entry zapcore.Entry) error {
		fmt.Print(customSystemLoggerFormat(entry))
		return nil
	}))

	defer SystemLog.Sync()
}

func InitLogger() {
	createSystemLogger()
	createDevLogger()
	createTraceLogger()
}
