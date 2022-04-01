/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/zalando/go-keyring"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Args:  cobra.ExactArgs(2),
	Short: "Get a password from the OS credential store",
	Long:  `Get a password from the OS credential store`,
	Run: func(cmd *cobra.Command, args []string) {
		service := args[0]
		user := args[1]

		// set password
		pwd, err := keyring.Get(service, user)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(pwd)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.SetUsageTemplate(fmt.Sprintln(`Usage:
  keyring get <service> <user>`))

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
