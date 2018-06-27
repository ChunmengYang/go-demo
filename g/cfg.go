package g

import (
	"github.com/json-iterator/go"
	"log"
	"os"
	"sync"

	"github.com/toolkits/file"
)

type HttpConfig struct {
	Enabled  bool   `json:"enabled"`
	Listen   string `json:"listen"`
}

type SocketConfig struct {
	Enabled  bool    `json:"enabled"`
	IP   	 string  `json:"ip"`
	Port   	 int 	 `json:"port"`
}

type GlobalConfig struct {
	Debug         bool             `json:"debug"`
	Logfile       string           `json:"logfile"`
	Hostname      string           `json:"hostname"`
	Http          *HttpConfig      `json:"http"`
	Socket        *SocketConfig    `json:"socket"`
}

var (
	ConfigFile  string
	config      *GlobalConfig
	lock        = new(sync.RWMutex)
	logger 		*log.Logger
)

func InitLog() {
	fileName := Config().Logfile
	logFile, err := os.Create(fileName)
	if err != nil {
		log.Fatalln("open file error !")
	}
	logger = log.New(logFile, "[Debug]", log.LstdFlags)
	log.Println("logging on", fileName)
}

func Logger() *log.Logger {
	lock.RLock()
	defer lock.RUnlock()
	return logger
}

func Config() *GlobalConfig {
	lock.RLock()
	defer lock.RUnlock()
	return config
}

func Hostname() (string, error) {
	hostname := Config().Hostname
	if hostname != "" {
		return hostname, nil
	}

	hostname, err := os.Hostname()
	if err != nil {
		log.Println("ERROR: os.Hostname() fail", err)
	}
	return hostname, err
}

func ParseConfig(cfg string) {
	if cfg == "" {
		log.Fatalln("use -c to specify configuration file")
	}

	if !file.IsExist(cfg) {
		log.Fatalln("config file:", cfg, "is not existent. maybe you need `mv cfg.example.json cfg.json`")
	}

	ConfigFile = cfg

	configContent, err := file.ToTrimString(cfg)
	if err != nil {
		log.Fatalln("read config file:", cfg, "fail:", err)
	}

	var c GlobalConfig
	err = jsoniter.Unmarshal([]byte(configContent), &c)
	if err != nil {
		log.Fatalln("parse config file:", cfg, "fail:", err)
	}

	lock.Lock()
	defer lock.Unlock()

	config = &c

	log.Println("read config file:", cfg, "successfully")
}
