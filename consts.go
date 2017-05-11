package kraken

import "time"

// TimeFormat defines the default format of all time related data.
const TimeFormat string = time.RFC3339

// Host adress of the server. Default to localhost, binding all IPs.
const Host string = ""

// Port of the server to listen to.
const Port int = 8000

// DefaultStore is the path for default storage
const DefaultStore string = "./"

// FileSuffix is the suffix for storage files
const FileSuffix string = ".kraken"
