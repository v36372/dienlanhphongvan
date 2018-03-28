package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	once             sync.Once
	rootConfigFile   string
	rootViper        = viper.New()
	customConfigFile string
	customViper      = viper.New()
	isHeroku         bool

	// rootCmd represents the base command when called without any subcommands
	rootCmd *cobra.Command
)

func Root() *cobra.Command {
	return rootCmd
}

func SetRunFunc(run func()) {
	rootCmd.Run = func(cmd *cobra.Command, args []string) {
		run()
	}
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func defaultFlags() {
	pflags := rootCmd.PersistentFlags()

	// app port
	pflags.Int("port", 9000, "app binding port")
	customViper.BindPFlag("app.port", pflags.Lookup("port"))

	// TODO: Add more default config here
}

func init() {
	cobra.OnInitialize(initCustomConfig)

	rootCmd = &cobra.Command{
		Use:   "CLI",
		Short: "Some help",
		// Uncomment the following line if your bare application
		// has an action associated with it:
		Run: func(cmd *cobra.Command, args []string) {},
	}
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	pflags := rootCmd.PersistentFlags()
	pflags.StringVar(&rootConfigFile, "rootConfig", "", "root config file, could be overrided (default is $GOPATH/src/dienlanhphongvan/config/config.yaml)")
	pflags.StringVar(&customConfigFile, "config", "config/config.yaml", "custom config file")
	pflags.BoolVar(&isHeroku, "heroku", false, "is deploy on heroku")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.

	defaultFlags()
}

func initCustomConfig() {
	if isHeroku {
		fmt.Println("deploted on herok")
		return
	}
	if customConfigFile != "" {
		customViper.SetConfigFile(customConfigFile)
		if err := customViper.ReadInConfig(); err != nil {
			fmt.Println("WARNING: file config/config.yaml not exist")
		}
	}
}

// initConfig reads in config file and ENV variables if set.
func initRootConfig() {
	if rootConfigFile != "" {
		// Use config file from the flag.
		rootViper.SetConfigFile(rootConfigFile)
	} else {
		envGoPath := os.Getenv("GOPATH")
		goPaths := filepath.SplitList(envGoPath)
		if len(goPaths) == 0 {
			panic("$GOPATH is not set")
		}
		for _, goPath := range goPaths {
			configDir := filepath.Join(goPath, "src", "dienlanhphongvan", "config")
			rootViper.AddConfigPath(configDir)
		}
		rootViper.SetConfigName("config")
	}
	/*
		// uncomment this if want to override config by env
		//  read in environment variables that match
		viper.SetEnvPrefix("LZ")
		viper.AutomaticEnv()
	*/

	// If a config file is found, read it in.
	if err := rootViper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func GetViper() *viper.Viper {
	fmt.Println("once 2")
	once.Do(func() {
		for _, key := range customViper.AllKeys() {
			rootViper.Set(key, customViper.Get(key))
		}
	})
	return rootViper
}

func IsOnHeroku() bool {
	return isHeroku
}
