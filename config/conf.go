package conf

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type (
	Config struct{}
)

func (c *Config) Load(fileName string) (cfg []byte) {
	var settingPath string
	var filePath string
	var sErr error
	flag.StringVar(&settingPath, "setting-path", "setting-path", "The config file path")
	flag.Parse()
	filePath = settingPath
	if !strings.HasSuffix(settingPath, "/") {
		filePath = settingPath + "/"
	}
	f, fErr := os.Open(filePath + fileName)
	if fErr != nil {
		log.Fatalln(fErr)
	}
	defer f.Close()
	cfg, sErr = ioutil.ReadAll(f)
	if sErr != nil {
		log.Fatalln(sErr)
	}
	return
}
