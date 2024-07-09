package config

import (
	_ "embed"
	"fmt"
	"os"
	"strconv"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap/zapcore"
)

// 系统配置，对应yml
// viper内置了mapstructure, yml文件用"-"区分单词, 转为驼峰方便

// 全局配置变量
var Conf = new(config)

type config struct {
	System   *SystemConfig `mapstructure:"system" json:"system"`
	Logs     *LogsConfig   `mapstructure:"logs" json:"logs"`
	Database *Database     `mapstructure:"database" json:"database"`
	Mysql    *MysqlConfig  `mapstructure:"mysql" json:"mysql"`
}

// 设置读取配置信息
func InitConfig() {
	workDir, err := os.Getwd()
	if err != nil {
		panic(fmt.Errorf("读取应用目录失败:%s", err))
	}
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/")
	// 读取配置信息
	err = viper.ReadInConfig()

	// 热更新配置
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		// 将读取的配置信息保存至全局变量Conf
		if err := viper.Unmarshal(Conf); err != nil {
			panic(fmt.Errorf("初始化配置文件失败:%s", err))
		}
	})

	if err != nil {
		panic(fmt.Errorf("读取配置文件失败:%s", err))
	}
	// 将读取的配置信息保存至全局变量Conf
	if err := viper.Unmarshal(Conf); err != nil {
		panic(fmt.Errorf("初始化配置文件失败:%s", err))
	}

	// 部分配合通过环境变量加载
	dbDriver := os.Getenv("DB_DRIVER")
	if dbDriver != "" {
		Conf.Database.Driver = dbDriver
	}
	mysqlHost := os.Getenv("MYSQL_HOST")
	if mysqlHost != "" {
		Conf.Mysql.Host = mysqlHost
	}
	mysqlUsername := os.Getenv("MYSQL_USERNAME")
	if mysqlUsername != "" {
		Conf.Mysql.Username = mysqlUsername
	}
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	if mysqlPassword != "" {
		Conf.Mysql.Password = mysqlPassword
	}
	mysqlDatabase := os.Getenv("MYSQL_DATABASE")
	if mysqlDatabase != "" {
		Conf.Mysql.Database = mysqlDatabase
	}
	mysqlPort := os.Getenv("MYSQL_PORT")
	if mysqlPort != "" {
		Conf.Mysql.Port, _ = strconv.Atoi(mysqlPort)
	}
}

type SystemConfig struct {
	Mode          string `mapstructure:"mode" json:"mode"`
	UrlPathPrefix string `mapstructure:"url-path-prefix" json:"urlPathPrefix"`
	Port          int    `mapstructure:"port" json:"port"`
	InitData      bool   `mapstructure:"init-data" json:"initData"`
}

type LogsConfig struct {
	Level      zapcore.Level `mapstructure:"level" json:"level"`
	Path       string        `mapstructure:"path" json:"path"`
	MaxSize    int           `mapstructure:"max-size" json:"maxSize"`
	MaxBackups int           `mapstructure:"max-backups" json:"maxBackups"`
	MaxAge     int           `mapstructure:"max-age" json:"maxAge"`
	Compress   bool          `mapstructure:"compress" json:"compress"`
}

type Database struct {
	Driver string `mapstructure:"driver" json:"driver"`
	Source string `mapstructure:"source" json:"source"`
}

type MysqlConfig struct {
	Username    string `mapstructure:"username" json:"username"`
	Password    string `mapstructure:"password" json:"password"`
	Database    string `mapstructure:"database" json:"database"`
	Host        string `mapstructure:"host" json:"host"`
	Port        int    `mapstructure:"port" json:"port"`
	Query       string `mapstructure:"query" json:"query"`
	LogMode     bool   `mapstructure:"log-mode" json:"logMode"`
	TablePrefix string `mapstructure:"table-prefix" json:"tablePrefix"`
	Charset     string `mapstructure:"charset" json:"charset"`
	Collation   string `mapstructure:"collation" json:"collation"`
}
