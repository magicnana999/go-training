package logger

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var Logger *zap.SugaredLogger

func getEncoder() zapcore.Encoder {
	//return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	//return zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig())

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000")
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	//encoderConfig.FunctionKey = "printMethodName"
	encoderConfig.ConsoleSeparator = " "

	return zapcore.NewConsoleEncoder(encoderConfig)

}

func getLogWriter() zapcore.WriteSyncer {

	lumberJackLogger := &lumberjack.Logger{
		Filename:   "log.log",
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   true,
	}
	return zapcore.NewMultiWriteSyncer(zapcore.AddSync(lumberJackLogger), zapcore.AddSync(os.Stdout))
}

func init() {
	Logger = _initLogger()
}
func _initLogger() *zap.SugaredLogger {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	//_logger := zap.New(core)
	_logger := zap.New(core, zap.AddCaller())

	l := _logger.Sugar()

	defer l.Sync()

	return l

}
