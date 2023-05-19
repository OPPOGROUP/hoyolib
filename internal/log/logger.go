package log

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"os"
	"path"
	"time"
)

var (
	logger zerolog.Logger
)

func Init() error {
	level, err := zerolog.ParseLevel(viper.GetString("log.level"))
	if err != nil {
		return err
	}
	zerolog.SetGlobalLevel(level)
	zerolog.TimeFieldFormat = "2006-01-02 15:04:05"

	dir := viper.GetString("log.path")
	if dir == "" {
		dir = "./log"
	}
	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}
	fileName := path.Join(dir, fmt.Sprintf("%s.log", time.Now().Format("2006-01-02")))
	log.Debug().Str("fileName", fileName)
	logFile, _ := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if viper.GetString("env") == "prod" {
		writer := zerolog.MultiLevelWriter(logFile)
		logger = zerolog.New(writer).Level(level).With().Timestamp().Logger()
		logger.s
	} else {

	}
}
