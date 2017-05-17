package kraken

import "time"

// ApplicationName to identify itself.
const ApplicationName string = "Kraken"

// ApplicationVersion current version.
const ApplicationVersion string = "v0.0.1"

// TimeFormat defines the default format of all time related data.
const TimeFormat string = time.RFC3339

// Host adress of the server. Default to localhost, binding all IPs.
const Host string = ""

// Port of the server to listen to.
const Port int = 8000

// DefaultStore is the path for default storage.
const DefaultStore string = "./"

// FileSuffix is the suffix for storage files.
const FileSuffix string = ".kraken"

// AutoWriteInterval is the duration of the interval for the auto persistent thread.
const AutoWriteInterval time.Duration = 10 * time.Second
