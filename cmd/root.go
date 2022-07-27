/*
Copyright © 2022 y16ra

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	cfg     Config
)

type Config struct {
	UserName string `mapstructure:"user_name"`
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-cobra",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// RunE: func(cmd *cobra.Command, args []string) error {
	// 	fmt.Println("root called")
	// 	return nil
	// },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-cobra.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "config", "Specify config file name. Looking up configs directory by this file name.")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigName(cfgFile)
	} else {
		viper.SetConfigName("config")
	}
	viper.AddConfigPath("configs/")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("invalid config file: %s \n", err)
		os.Exit(1)
	}
	if err := viper.Unmarshal(&cfg); err != nil {
		fmt.Printf("unmarshal failed : %s \n", err)
		os.Exit(1)
	}
}
