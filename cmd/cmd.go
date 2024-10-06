package cmd

import (
	"fmt"
	"os"

  alice "github.com/ohhfishal/alice/api/v1"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Config struct {
  Config alice.Config 
}

var c Config

func init() {
	rootCmd.PersistentFlags().StringVarP(&c.Config.Filepath, "file", "f", "test.txt", "task file")
  viper.BindPFlag("file", rootCmd.PersistentFlags().Lookup("file"))
}

var rootCmd = &cobra.Command{
	Use:   "alice",
	Short: "alice is a simple task/event tracker",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello world")
		fmt.Println(c.Config)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
