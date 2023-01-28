package zap

import (
	"gopkg.in/natefinch/lumberjack.v2"
	"testing"

	"github.com/AarenWang/go-log/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type testWriteSyncer struct {
	output []string
}

func (x *testWriteSyncer) Write(p []byte) (n int, err error) {
	x.output = append(x.output, string(p))
	return len(p), nil
}

func (x *testWriteSyncer) Sync() error {
	return nil
}

func RollFileWriterSyncer() zapcore.WriteSyncer {
	hook := &lumberjack.Logger{
		Filename:   "./logs/roll_app" + ".log",
		MaxSize:    1,    //日志最大的大小（M）
		MaxBackups: 7,    //备份个数
		MaxAge:     7,    //最大保存天数（day）
		Compress:   true, //是否压缩
		LocalTime:  false,
	}

	w := zapcore.AddSync(hook)
	return w
}

func getEncode() zapcore.EncoderConfig {
	encoderCfg := zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		NameKey:        "logger",
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
	}
	return encoderCfg
}

// 滚动输出 归档
func TestLogger(t *testing.T) {
	//syncer := &testWriteSyncer{}
	syncer := RollFileWriterSyncer()
	encoderCfg := getEncode()
	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoderCfg), syncer, zap.DebugLevel)
	zlogger := zap.New(core).WithOptions()
	logger := NewLogger(zlogger)

	defer func() { _ = logger.Close() }()

	zlog := log.NewHelper(logger)

	//只打印键值对
	zlog.Debugw("log", "debug")
	zlog.Infow("log", "info")
	zlog.Warnw("log", "warn")
	zlog.Errorw("log", "error")
	zlog.Errorw("log", "error", "except warn")

	//格式化字符串
	zlog.Infof("This is Info Level Log username=%s", "zhangsan")
	zlog.Warnf("This is Warn Level Log username=%s", "zhangsan")

}

func TestLoggerField(t *testing.T) {
	syncer := RollFileWriterSyncer()
	encoderCfg := getEncode()
	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoderCfg), syncer, zap.DebugLevel)
	//zlogger := zap.New(core).WithOptions()
	zlogger := zap.New(core).WithOptions(zap.AddCaller(), zap.WithCaller(true)) // zap.AddCaller() 不起作用
	logger := NewLogger(zlogger)

	// 编译不通过
	//logger = log.With(logger, "ts", log.DefaultTimestamp, "caller", log.DefaultCaller)

	defer func() { _ = logger.Close() }()

	helperlog := log.NewHelper(logger)
	// 编译不通过
	//helperlog = log.With(helperlog, "ts", log.DefaultTimestamp, "caller", log.DefaultCaller)

	//只打印键值对
	helperlog.Infow("log", "info")

	//格式化字符串
	helperlog.Infof("This is Info Level Log username=%s", "zhangsan")
}
