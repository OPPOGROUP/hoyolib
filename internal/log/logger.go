package log

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"os"
	"path"
	"strings"
	"time"
)

var (
	logger zerolog.Logger
)

func GetLogger() zerolog.Logger {
	return logger
}

func Debug() *zerolog.Event {
	return logger.Debug()
}

func Info() *zerolog.Event {
	return logger.Info()
}

func Warn() *zerolog.Event {
	return logger.Warn()
}

func Error() *zerolog.Event {
	return logger.Error().Stack()
}

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
	var writer zerolog.LevelWriter
	if viper.GetString("env") == "prod" {
		writer = zerolog.MultiLevelWriter(logFile)
	} else {
		consoleWriter := zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: "2006-01-02 15:04:05",
		}
		consoleWriter.FormatLevel = func(i interface{}) string {
			return strings.ToUpper(fmt.Sprintf("|%-6s|", i))
		}
		consoleWriter.FormatMessage = func(i interface{}) string {
			return fmt.Sprintf("%s", i)
		}
		consoleWriter.FormatFieldName = func(i interface{}) string {
			return fmt.Sprintf("[ %s: ", i)
		}
		consoleWriter.FormatFieldValue = func(i interface{}) string {
			return fmt.Sprintf("%s ]", i)
		}
		writer = zerolog.MultiLevelWriter(logFile, consoleWriter)
	}
	logger = zerolog.New(writer).Level(level).With().Timestamp().Caller().Logger()
	return nil
}
