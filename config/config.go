package config

import (
	"dienlanhphongvan/cmd"
	"fmt"
	"os"
	"strconv"
	"sync"
)

var (
	conf Config
	once sync.Once
)

type Config struct {
	App         App
	CookieToken CookieToken
	PostgreSQL  PostgreSQL
	Log         Log
	Resource    Resource
	Imgx        Service
}

type CookieToken struct {
	HashKey  string
	BlockKey string
}

type Resource struct {
	RootDir string
}

type Service struct {
	Address string
}

type App struct {
	Host      string
	Port      int
	Debug     bool
	Whitelist []string
}

type Log struct {
	Prefix     string
	Dir        string
	LevelDebug bool
}

type PostgreSQL struct {
	Username     string
	Password     string
	Host         string
	Port         int
	Db           string
	Debug        bool
	MaxIdleConns int
	MaxOpenConns int
}

func init() {
	// Init CLI commands
	cmd.Root().Use = "bin/dienlanhphongvan --config <Config path>"
	cmd.Root().Short = "dienlanhphongvan - Provide API for dienlanhphongvan"
	cmd.Root().Long = "dienlanhphongvan"

	cmd.SetRunFunc(load)
}

func loadFromOS(conf *Config) {
	conf.App.Debug, _ = strconv.ParseBool(os.Getenv("debug"))
	conf.App.Host = os.Getenv("host")
	conf.App.Port, _ = strconv.Atoi(os.Getenv("PORT"))

	conf.CookieToken.BlockKey = os.Getenv("blockkey")
	conf.CookieToken.HashKey = os.Getenv("hashkey")

	conf.PostgreSQL.Username = os.Getenv("dbusername")
	conf.PostgreSQL.Password = os.Getenv("dbpassword")
	conf.PostgreSQL.Host = os.Getenv("dbhost")
	conf.PostgreSQL.Port, _ = strconv.Atoi(os.Getenv("dbport"))
	conf.PostgreSQL.Db = os.Getenv("dbname")
	conf.PostgreSQL.Debug, _ = strconv.ParseBool(os.Getenv("dbdebug"))
	conf.PostgreSQL.MaxIdleConns, _ = strconv.Atoi(os.Getenv("dbmaxidleconns"))
	conf.PostgreSQL.MaxOpenConns, _ = strconv.Atoi(os.Getenv("dbmaxopenconns"))

	conf.Log.Dir = os.Getenv("logdir")
	conf.Log.LevelDebug, _ = strconv.ParseBool(os.Getenv("loglevel"))
	conf.Log.Prefix = os.Getenv("logprefix")

	conf.Resource.RootDir = os.Getenv("resource")

	conf.Imgx.Address = os.Getenv("imgxaddress")

}

func load() {
	once.Do(func() {
		if cmd.IsOnHeroku() {
			loadFromOS(&conf)
			return
		}
		if err := cmd.GetViper().Unmarshal(&conf); err != nil {
			fmt.Println("load viper fail")
		}
	})
}

func Load() {
	load()
}

func Get() Config {
	load()
	return conf
}

func GetPostgreSQL() PostgreSQL { return conf.GetPostgreSQL() }
func (c Config) GetPostgreSQL() PostgreSQL {
	return c.PostgreSQL
}
