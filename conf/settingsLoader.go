package conf

import (
	"github.com/LaMonF/FDJ_SLACK/log"
	"github.com/LaMonF/FDJ_SLACK/model"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

const SETTING_FILE  = "settings.yml"

var Settings = loadSettings()

func loadSettings() model.Settings {
	s := model.Settings{}
	stringFile, err := ioutil.ReadFile(SETTING_FILE)
	if err != nil {
		log.Error("cannot READ setting file: " +SETTING_FILE,  err)
		os.Exit(1)
	}
	err = yaml.Unmarshal([]byte(stringFile), &s)
	if err != nil {
		log.Error("cannot LOAD setting file: " +SETTING_FILE,  err)
		os.Exit(1)
	}

	log.Info("Configuration : \n" + s.String())
	return s
}
