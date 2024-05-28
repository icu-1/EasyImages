package main

import (
	"gin-template/internal/common"
	"gin-template/logger"
	"gin-template/vars"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	port     int
	logLevel = "info"
	logPath  = "log"

	cmd = &cobra.Command{
		Short:   "GIN 模版",
		Long:    "项目地址：https://github.com/xxx/gin-template",
		Use:     vars.Project,
		Version: vars.Version,
		Run: func(cmd *cobra.Command, args []string) {
			common.InitCommon()
			logger.InitLogger(logPath, LogLevel())
			// TODO -
		},
	}
)

func main() {
	cmd.PersistentFlags().IntVar(&port, "port", 8080, "服务端口 port")
	cmd.PersistentFlags().StringVar(&logLevel, "log", logLevel, "日志级别: trace|debug|info|warn|error")
	cmd.PersistentFlags().StringVar(&logPath, "log-path", logPath, "日志路径")
	_ = cmd.Execute()
}

func LogLevel() logrus.Level {
	switch logLevel {
	case "trace":
		return logrus.TraceLevel
	case "debug":
		return logrus.DebugLevel
	case "warn":
		return logrus.WarnLevel
	case "error":
		return logrus.ErrorLevel
	default:
		return logrus.InfoLevel
	}
}
