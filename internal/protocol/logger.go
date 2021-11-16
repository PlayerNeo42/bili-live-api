package protocol

import (
	"github.com/sirupsen/logrus"
	"os"
)

var log *logrus.Logger

func InitLogger() {
	//var err error
	log = logrus.New()
	log.SetReportCaller(true)
	log.SetLevel(logrus.InfoLevel)
	//f, err := os.Create("server.log")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.SetOutput(f)
	log.SetOutput(os.Stdout)
	log.SetFormatter(&logrus.JSONFormatter{
		DisableHTMLEscape: true,
	})
}

func GetLogger() *logrus.Logger {
	if log == nil {
		InitLogger()
	}
	return log
}
