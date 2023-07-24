/**
 * @Create on : 2023/4/17
 * @Author: sunnyh
 * @Des:
 */

package settings

import (
	"fmt"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Name    string `mapstructure:"name"`
	Mode    string `mapstructure:"mode"`
	Version string `mapstructure:"version"`
	Port    int    `mapstructure:"port"`

	StaticUrlPrefix string `mapstructure:"static_url_prefix"`
	StaticRoot      string `mapstructure:"static_root"`

	*LogConfig `mapstructure:"log"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

var Conf Config

func Init() error {
	viper.SetConfigName("config")           // name of config file (without extension)
	viper.SetConfigType("yaml")             // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("/etc/arena-app/")  // path to look for the config file in
	viper.AddConfigPath("$HOME/.arena-app") // call multiple times to add many search paths
	viper.AddConfigPath(".")                // optionally look for config in the working directory

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	viper.SetEnvPrefix("awa") // Arena Web App
	viper.AutomaticEnv()
	viper.AllowEmptyEnv(true)

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		fmt.Printf("fatal error config file: %s \n", err)
		return err
	}

	if err := viper.Unmarshal(&Conf); err != nil {
		fmt.Printf("unable to decode into struct, %v \n", err)
	}

	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		if err := viper.Unmarshal(&Conf); err != nil {
			fmt.Printf("unable to decode into struct, %v \n", err)
		}
	})

	viper.WatchConfig()

	return err
}
