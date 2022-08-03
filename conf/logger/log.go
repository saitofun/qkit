package logger

import (
	"os"
	"strings"

	"github.com/saitofun/qkit/conf/log"
	"github.com/sirupsen/logrus"
)

var OutputAlways = &logrus.Logger{}

type Log struct {
	ReportCaller bool `env:""`
	Name         string
	Level        string `env:""`
	Output       log.Logger
	Format       string
	init         bool
}

func (log *Log) SetDefault() {
	if log.Name == "" {
		log.Name = os.Getenv("PROJECT_NAME")
	}
	log.ReportCaller = false

	if log.Level == "" {
		log.Level = "DEBUG"
	}

	if log.Format == "" {
		log.Format = "json"
	}
}

func (log *Log) Init() {
	if !log.init {
		log.Create()
		log.init = true
	}
}

func (log *Log) Create() {
	if log.Format == "json" {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	} else {
		logrus.SetFormatter(&logrus.TextFormatter{
			ForceColors: true,
		})
	}

	logrus.SetLevel(getLogLevel(log.Level))
	logrus.SetReportCaller(log.ReportCaller)
	// TODO add hook with goid meta
	// logrus.AddHook(goid.Default)
	logrus.AddHook(NewProjectHook(log.Name))
	// TODO export trace & span, such as zipkin ...

	logrus.SetOutput(os.Stdout)
}

func getLogLevel(l string) logrus.Level {
	level, err := logrus.ParseLevel(strings.ToLower(l))
	if err == nil {
		return level
	}
	return logrus.InfoLevel
}

func NewProjectHook(name string) *ProjectHook {
	return &ProjectHook{
		Name: name,
	}
}

type ProjectHook struct {
	Name string
}

func (hook *ProjectHook) Fire(entry *logrus.Entry) error {
	entry.Data["project"] = hook.Name
	return nil
}

func (hook *ProjectHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
