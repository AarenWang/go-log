module github.com/AarenWang/go-log/contrib/log/zap/v2

go 1.19

require (
	github.com/go-kratos/kratos/v2 v2.5.4
	github.com/natefinch/lumberjack v2.0.0+incompatible
	go.uber.org/zap v1.23.0
)

require (
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
)

replace github.com/AarenWang/go-log => ../../../
