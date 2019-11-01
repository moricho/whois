package logger

import (
	"time"

	"github.com/moricho/whois/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

var Logger = newLogger()

func Shutdown() error {
	return Logger.Sync()
}

func newLogger() *zap.Logger {
	zapCfg := zap.NewProductionConfig()
	encCfg := zapCfg.EncoderConfig
	encCfg.EncodeTime = jstTimeEncoder
	enc := zapcore.NewJSONEncoder(encCfg)

	ws := zapcore.AddSync(
		&lumberjack.Logger{
			Filename:   viper.GetString(config.LogPath),
			MaxSize:    viper.GetInt(config.LogRotateMaxSize), // MB
			MaxBackups: viper.GetInt(config.LogRotateMaxBackups),
			MaxAge:     viper.GetInt(config.LogRotateMaxDays),
		},
	)

	return zap.New(
		zapcore.NewCore(enc, ws, zapCfg.Level),
	)
}

func jstTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	const layout = "2006-01-02 15:04:05"
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	enc.AppendString(t.In(jst).Format(layout))
}
