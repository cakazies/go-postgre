package cmd

import (
	"fmt"
	"log"

	"github.com/cakazies/go-postgre/routes"
	"github.com/cakazies/go-postgre/utils"
	"github.com/spf13/cobra"
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
	cobra.OnInitialize(splash)
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

// Execute from Cobra Firsttime
func Execute() {
	err := rootCmd.Execute()
	utils.FailError(err, "Error Execute RootCMD")
}
