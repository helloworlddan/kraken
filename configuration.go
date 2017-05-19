package kraken

import "time"

// Configuration holds the entire config.
type Configuration struct {
	applicationName     string
	applicationVersion  string
	timeFormat          string
	host                string
	port                int
	defaultStore        string
	fileSuffix          string
	autoWriteInterval   time.Duration
	strictSlashesInURLs bool
}

// ApplicationName gets this Configuration's applicationName
func (c *Configuration) ApplicationName() string {
	return c.applicationName
}

// ApplicationVersion gets this Configuration's applicationVersion
func (c *Configuration) ApplicationVersion() string {
	return c.applicationVersion
}

// TimeFormat gets this Configuration's timeFormat
func (c *Configuration) TimeFormat() string {
	return c.timeFormat
}

// Host gets this Configuration's host
func (c *Configuration) Host() string {
	return c.host
}

// Port gets this Configuration's port
func (c *Configuration) Port() int {
	return c.port
}

// DefaultStore gets this Configuration's defaultStore
func (c *Configuration) DefaultStore() string {
	return c.defaultStore
}

// FileSuffix gets this Configuration's fileSuffix
func (c *Configuration) FileSuffix() string {
	return c.fileSuffix
}

// AutoWriteInterval gets this Configuration's autoWriteInterval
func (c *Configuration) AutoWriteInterval() time.Duration {
	return c.autoWriteInterval
}

// StrictSlashesInURLs gets this Configuration's strictSlashesInURLs
func (c *Configuration) StrictSlashesInURLs() bool {
	return c.strictSlashesInURLs
}

// DefaultConfiguration of the application.
func DefaultConfiguration() *Configuration {
	return &Configuration{
		applicationName:     "Kraken",
		applicationVersion:  "v0.0.1",
		timeFormat:          time.RFC3339,
		host:                "",
		port:                8000,
		defaultStore:        "./",
		fileSuffix:          ".kraken",
		autoWriteInterval:   time.Second * 10,
		strictSlashesInURLs: true,
	}
}
