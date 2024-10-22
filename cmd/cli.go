/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/pedrojpx/hexagonal-example/adapters/cli"
	"github.com/spf13/cobra"
)

var action string
var productId string
var productName string
var productPrice float64

// cliCmd represents the cli command
var cliCmd = &cobra.Command{
	Use:   "cli",
	Short: "creates, gets, enables or disables a product",
	Long: `creates, gets, enables or disables a product
	
available values for "action" flag are:
create, enable, disable and get`,
	Run: func(cmd *cobra.Command, args []string) {
		res, err := cli.Run(prodcutService, action, productId, productName, productPrice)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(res)
	},
}

func init() {
	rootCmd.AddCommand(cliCmd)
	cliCmd.Flags().StringVarP(&action, "action", "a", "enable", "Enable/Disable a product")
	cliCmd.Flags().StringVarP(&productId, "id", "i", "", "product id")
	cliCmd.Flags().StringVarP(&productName, "productName", "n", "", "product name")
	cliCmd.Flags().Float64VarP(&productPrice, "productPrice", "p", 0.0, "product name")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cliCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cliCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
