package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start an service",
	Long:  "Start an service",
	Example: "als start --config als.yml",
	Run: func(cmd *cobra.Command, args []string) {
		if err := cmd.Usage(); err != nil {
			if _, err := fmt.Fprint(os.Stderr, "execption"); err != nil {
				os.Exit(0)
			}
		}
	},
}

var conf string

func init() {
	rootCmd.AddCommand(startCmd)
	cobra.OnInitialize(initConfig)
	startCmd.PersistentFlags().StringVarP(&conf, "config", "c", "als.yml", "config")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if conf != "" {
		// Use config file from the flag.
		viper.SetConfigFile(conf)
	} else {
		// Search config in home directory with name ".als" (without extension).
		viper.AddConfigPath("/etc/als/als.yml")
		viper.SetConfigName("als")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
