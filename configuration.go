package kraken

import (
	"io/ioutil"
	"log"
	"time"

	"gopkg.in/yaml.v2"
)

// ConfigurationPath default path to load configuration.
const ConfigurationPath = "./config.yaml"

// Configuration holds the entire config.
type Configuration struct {
	ApplicationName     string
	ApplicationVersion  string
	TimeFormat          string
	Host                string
	Port                int
	DefaultStore        string
	FileSuffix          string
	AutoWriteInterval   time.Duration
	StrictSlashesInURLs bool
	OutputFormat        string
}

func loadFromDisk() (*Configuration, error) {
	data, err := ioutil.ReadFile(ConfigurationPath)
	if err != nil {
		return nil, err
	}
	conf := defaultConfiguration()
	err = yaml.Unmarshal([]byte(data), conf)
	if err != nil {
		return nil, err
	}
	log.Println("Loaded configuration from " + ConfigurationPath)
	return conf, nil
}

// UseConfiguration returns the currently valid configuration.
func UseConfiguration() *Configuration {
	conf, err := loadFromDisk()
	if err != nil {
		return defaultConfiguration()
	}
	return conf
}

func defaultConfiguration() *Configuration {
	return &Configuration{
		ApplicationName:     "Kraken",
		ApplicationVersion:  "v0.0.1",
		TimeFormat:          time.RFC3339,
		Host:                "",
		Port:                8000,
		DefaultStore:        "./",
		FileSuffix:          ".kraken",
		AutoWriteInterval:   time.Second * 10,
		StrictSlashesInURLs: true,
		OutputFormat:        "JSON",
	}
}
