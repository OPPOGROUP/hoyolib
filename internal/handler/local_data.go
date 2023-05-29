package handler

import (
	"encoding/json"
	"github.com/OPPOGROUP/hoyolib/internal/cte"
	"github.com/OPPOGROUP/hoyolib/internal/log"
	"github.com/OPPOGROUP/hoyolib/internal/user"
	"github.com/spf13/viper"
	"os"
	"path"
)

var (
	m         = make(map[int64]*user.Info)
	uid int64 = 100000
)

func GetUserData() map[int64]*user.Info {
	return m
}

func LoadSavedUsers() {
	dir := viper.GetString("data.path")
	if dir == "" {
		dir = "./data"
	}
	filename := path.Join(dir, cte.UserDataFile)
	userBytes, err := os.ReadFile(filename)
	if err != nil {
		log.Warn().Err(err).Msg("Load local user data file failed")
		return
	}
	_ = json.Unmarshal(userBytes, &m)
}

func saveUser() error {
	dir := viper.GetString("data.path")
	if dir == "" {
		dir = "./data"
	}
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}
	filename := path.Join(dir, cte.UserDataFile)
	userJson, _ := json.Marshal(m)
	err = os.WriteFile(filename, userJson, os.ModePerm)
	if err != nil {
		log.Error().Err(err).Msg("Save user failed")
		return err
	}
	return nil
}
