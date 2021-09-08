package config

import (
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	defCfg      map[string]string
	initialized = false
)

// initialize this configuration
func initialize() {
	viper.SetEnvPrefix("MW_TEST")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	defCfg = make(map[string]string)

	defCfg["server.env"] = "PRODUCTION" // DEVELOPMENT | STAGING | PRODUCTION
	defCfg["server.log.level"] = "info" // valid values are trace, debug, info, warn, error, fatal

	//Configuration for log file
	defCfg["log.type"] = "CMD" // FILE, CMD
	defCfg["log.path"] = "storage/logs"
	defCfg["log.max.age"] = "7" //days

	//Configuration api service
	defCfg["server.host"] = "127.0.0.1"
	defCfg["server.port"] = "8080"
	defCfg["server.context.timeout"] = "30" // seconds

	//Configuration db
	defCfg["db.type"] = "mysql"
	defCfg["db.user"] = "backend"
	defCfg["db.password"] = "backend"
	defCfg["db.database"] = "backend"
	defCfg["db.host"] = "127.0.0.1"
	defCfg["db.port"] = "3306"

	// time
	defCfg["time.default"] = "02 Jan 70 00:00 WIB" // RFC822 --> 1970-01-02 00:00:00

	for k := range defCfg {
		err := viper.BindEnv(k)
		if err != nil {
			log.Errorf("Failed to bind env \"%s\" into configuration. Got %s", k, err)
		}
	}

	initialized = true
}

// SetConfig put configuration key value
func SetConfig(key, value string) {
	viper.Set(key, value)
}

// Get fetch configuration as string value
func Get(key string) string {
	if !initialized {
		initialize()
	}
	ret := viper.GetString(key)
	if len(ret) == 0 {
		if ret, ok := defCfg[key]; ok {
			return ret
		}
		log.Debugf("%s config key not found", key)
	}
	return ret
}

// GetBoolean fetch configuration as boolean value
func GetBoolean(key string) bool {
	if len(Get(key)) == 0 {
		return false
	}
	b, err := strconv.ParseBool(Get(key))
	if err != nil {
		panic(err)
	}
	return b
}

// GetInt fetch configuration as integer value
func GetInt(key string) int {
	if len(Get(key)) == 0 {
		return 0
	}
	i, err := strconv.ParseInt(Get(key), 10, 64)
	if err != nil {
		panic(err)
	}
	return int(i)
}

// GetFloat fetch configuration as float value
func GetFloat(key string) float64 {
	if len(Get(key)) == 0 {
		return 0
	}
	f, err := strconv.ParseFloat(Get(key), 64)
	if err != nil {
		panic(err)
	}
	return f
}

// Set configuration key value
func Set(key, value string) {
	defCfg[key] = value
}
