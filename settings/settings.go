package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Conf全局变量，用来保存程序的所有配置信息
var Conf = new(AppConfig)

type AppConfig struct {
	Name         string `mapstructure:"name"`
	Mode         string `mapstructure:"mode"`
	Version      string `mapstructure:"version"`
	StartTime    string `mapstructure:"start_time"`
	MachineID    int64 `mapstructure:"machine_id"`
	Port         int    `mapstructure:"port"`
	*LogConfig   `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}
type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}
type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DbName       string `mapstructure:"db_name"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}
type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	DB       int    `mapstructure:"db"`
	Password string `mapstructure:"password"`
	PoolSize int    `mapstructure:"pool_size"`
}

func Init() (err error) {
	// 获取方式一
	//viper.SetConfigFile("config.json") // 指定本地配置文件可包含路径(json格式)
	viper.SetConfigFile("./config.yaml") // 指定本地配置文件可包含路径（yaml格式）

	// 获取方式二
	//viper.SetConfigName("config")        // 指定配置文件名称（不需要带后缀，如果获取本地则直接通过名字查找第一个）
	//viper.SetConfigType("yaml")        // 指定配置文件类型（专用于从远程获取配置信息时指定配置文件类型）
	//viper.AddConfigPath(".")      // 指定配置文件的路径（这里是相对路径，跟SetConfigName配合使用）
	//viper.AddConfigPath("./conf") // 指定配置文件的路径（可以有多个，代表如果当前目录找不到，则去当前目录conf目录去找）

	err = viper.ReadInConfig() // 读取配置文件
	if err != nil {
		//读取配置文件失败
		fmt.Printf("viper.ReadInConfig() failed，err:%v\n", err)
		return
	}
	// 把读取到的配置信息反序列化到Conf变量中
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal failed, err: %v\n", err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		//配置文件发生变更之后调用的回调函数
		fmt.Println("配置文件修改好了")
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper.Unmarshal failed, err: %v\n", err)
		}
	})
	return
}
