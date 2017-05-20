package kraken

import "time"

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

// DefaultConfiguration of the application.
func DefaultConfiguration() *Configuration {
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
