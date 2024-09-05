package conf

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"path/filepath"
	"time"
)

func NewLogger() *zap.SugaredLogger {
	logLevel := zap.DebugLevel
	if viper.GetString("build.type") == "Release" {
		logLevel = zap.InfoLevel
	}
	core := zapcore.NewCore(getEncoder(), getWriterSyncer(), logLevel)

	logger := zap.New(core).Sugar()
	logger.Infoln("========== Blazing fast, structured, leveled logging in Go. ==========")
	return logger
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder // 大写日志级别 INFO, WARN, DEBUG, ...
	encoderConfig.TimeKey = "time"                          // ts (timestamp) => time
	encoderConfig.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(t.Local().Format(time.DateTime)) // time.DateTime = "2006-01-02 15:04:05"
	}
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getWriterSyncer() zapcore.WriteSyncer {
	sep := string(filepath.Separator)
	rootDir, _ := os.Getwd()
	// logFile := fmt.Sprintf("./log/%s.log", time.Now().Format(time.DateOnly))
	logFile := rootDir + sep + "log" + sep + time.Now().Format(time.DateOnly) + ".log"
	log.Printf("logFile = %s\n", logFile)
	//! O_CREATE 如果文件不存在，则创建文件
	//! O_RDONLY 只读 | O_WRONLY 只写 | O_RDWR 读写
	//! O_APPEND 追加 | O_TRUNC 重写
	fp, _ := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	return zapcore.AddSync(fp /* fp 实现了 io.Writer 接口 */)
}
