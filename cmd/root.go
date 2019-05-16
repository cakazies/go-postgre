package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"

	conf "github.com/local/go-postgre/models"
	"github.com/spf13/cobra"
)

var (
	cfgFile string
)

var rootCmd = &cobra.Command{
	Use:   "GO-POSTGRESQL",
	Short: "Hugo is a very fast static site generator",
	Long: `A fast and Flexible Static Site Generator built width
			love by spf13 and friends in Go. Complete Documentation
			is avaible at http://hugo.spf13.com`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("go-postgresql is avaible running")
		Route()
	},
}

func init() {
	cobra.OnInitialize(initViper, conf.Connect)
}

func initViper() {
	viper.SetConfigFile("toml")
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath("./configs")
		viper.SetConfigName("config")
	}
	//
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error Viper config", err)
	}
	fmt.Println("Using Config File: ", viper.ConfigFileUsed())
}
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
