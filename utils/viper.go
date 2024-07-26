package utils

import (
	"fmt"
	_viper "github.com/spf13/viper"
	"log"
	"sync"
)

type Viper interface {
	ConfigInterface
}

type viper struct {
	viper *_viper.Viper
	*sync.Mutex
}

func (v *viper) GetInt(key string) int         { return v.viper.GetInt(ConfigRootKey + key) }
func (v *viper) GetInt64(key string) int64     { return v.viper.GetInt64(ConfigRootKey + key) }
func (v *viper) GetFloat64(key string) float64 { return v.viper.GetFloat64(ConfigRootKey + key) }
func (v *viper) GetBool(key string) bool       { return v.viper.GetBool(ConfigRootKey + key) }
func (v *viper) GetString(key string) string   { return v.viper.GetString(ConfigRootKey + key) }
func (v *viper) GetStringSlice(key string) []string {
	return v.viper.GetStringSlice(ConfigRootKey + key)
}

var Config Viper

func InitConfig() error {
	v := _viper.New()
	v.SetConfigName(".env")
	v.SetConfigType("json")
	v.AddConfigPath(".")

	if err := v.ReadInConfig(); err != nil {
		v = _viper.New()
		v.AutomaticEnv()
		v.Set("test.key", "test.value")
		log.Printf("Test Key Value from v: %s", v.GetString("test.key"))
		log.Printf("Test Key Value from viper: %s", _viper.GetString("test.key"))
		log.Printf("Config logger path from v: %s", v.GetString(ConfigRootKey+LoggerPath))
		log.Printf("Config logger path from viper: %s", _viper.GetString(ConfigRootKey+LoggerPath))
		fmt.Println("Config file not found; falling back to environment variables.")
	}

	Config = &viper{
		viper: v,
		Mutex: &sync.Mutex{},
	}

	return nil
}
