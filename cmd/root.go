/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"database/sql"
	"os"

	persistence "github.com/pedrojpx/hexagonal-example/adapters/db"
	"github.com/pedrojpx/hexagonal-example/application"
	"github.com/spf13/cobra"
)

var db, _ = sql.Open("sqlite3", "db.sqlite")
var productDb = persistence.NewProductDb(db)
var prodcutService = application.NewProductService(productDb)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "hexagonal-example",
	Short: "A simple application meant to demonstrate an use case of hexagonal architecture",
	Long: `
Hexagonal architecture strives to maintain the core of an application uncoupled from 
the methods of acessing the core application
	
in this case specifically, this cli tool is able to create, view, enable and disable a
product using the same core code as a different tool that does the same via a web server,
thus ensuring the same behavior accross different methods`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.hexagonal-example.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
