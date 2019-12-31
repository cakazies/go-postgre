package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/viper"

	conf "github.com/cakazies/go-postgre/application/models"
	"github.com/cakazies/go-postgre/routes"
	"github.com/cakazies/go-postgre/utils"
	"github.com/spf13/cobra"
)

var (
	cfgFile string
)

var rootCmd = &cobra.Command{
	Use:   "GO-POSTGRESQL",
	Short: "Tutorial golang in postgresql",
	Long:  `tutorial golang in postgresql and some plugins`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("go-postgre is avaible running")
		routes.Route()
	},
}

func init() {
	cobra.OnInitialize(splash, InitViper, conf.Connect)
}

// Opened
func splash() {
	fmt.Println(`
	______            ____             __                
	/ ____/___        / __ \____  _____/ /_____ _________ 
   / / __/ __ \______/ /_/ / __ \/ ___/ __/ __ / ___/ _ \
  / /_/ / /_/ /_____/ ____/ /_/ (__  ) /_/ /_/ / /  /  __/
  \____/\____/     /_/    \____/____/\__/\__, /_/   \___/ 
										/____/            
  `)
	// http://patorjk.com
}

// InitViper from file toml
func InitViper() {
	viper.SetConfigFile("toml")
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath("./configs")
		viper.SetConfigName("config")
	}
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	utils.FailError(err, "Error Viper config")
	// log.Println("Using Config File: ", viper.ConfigFileUsed())
}

// Execute from Cobra Firsttime
func Execute() {
	err := rootCmd.Execute()
	utils.FailError(err, "Error Execute RootCMD")
}
