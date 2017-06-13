package core

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
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

func (c *Configuration) loadFromDisk() (err error) {
	if c == nil {
		c = defaultConfiguration()
	}
	data, err := ioutil.ReadFile(ConfigurationPath)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal([]byte(data), c)
	if err != nil {
		return err
	}
	log.Println("Loaded configuration from " + ConfigurationPath)
	return nil
}

func (c *Configuration) loadFromEnvironment() (err error) {
	if c == nil {
		c = defaultConfiguration()
	}

	if key := os.Getenv("KRAKEN_APPLICATIONNAME"); key != "" {
		c.ApplicationName = key
	}
	if key := os.Getenv("KRAKEN_APPLICATIONVERSION"); key != "" {
		c.ApplicationVersion = key
	}
	if key := os.Getenv("KRAKEN_TIMEFORMAT"); key != "" {
		c.TimeFormat = key
	}
	c.Host = os.Getenv("KRAKEN_HOST")
	if key := os.Getenv("KRAKEN_PORT"); key != "" {
		i, err := strconv.Atoi(key)
		if err != nil {
			return err
		}
		c.Port = i
	}
	if key := os.Getenv("KRAKEN_DEFAULTSTORE"); key != "" {
		c.DefaultStore = key
	}
	if key := os.Getenv("KRAKEN_FILESUFFIX"); key != "" {
		c.FileSuffix = key
	}
	if key := os.Getenv("KRAKEN_AUTOWRITEINTERVAL"); key != "" {
		dur, err := time.ParseDuration(key)
		if err != nil {
			return err
		}
		c.AutoWriteInterval = dur
	}
	if key := os.Getenv("KRAKEN_STRICTSLASHESINURLS"); key != "" {
		b, err := strconv.ParseBool(key)
		if err != nil {
			return err
		}
		c.StrictSlashesInURLs = b
	}
	if key := os.Getenv("KRAKEN_OUTPUTFORMAT"); key != "" {
		c.OutputFormat = key
	}

	return nil
}

// UseConfiguration returns the currently valid configuration.
func UseConfiguration() (conf *Configuration) {
	conf = defaultConfiguration()
	err := conf.loadFromDisk()
	if err != nil {
		log.Println(err)
	}
	err = conf.loadFromEnvironment()
	if err != nil {
		log.Println(err)
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
