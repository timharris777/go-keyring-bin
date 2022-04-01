/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	utils "github.com/timharris777/go-keyring-bin/helper"
	"github.com/zalando/go-keyring"
)

// delCmd represents the del command
var delCmd = &cobra.Command{
	Use:   "del",
	Args:  cobra.ExactArgs(2),
	Short: "Delete a password from the OS credential store",
	Long:  `Delete a password from the OS credential store`,
	Run: func(cmd *cobra.Command, args []string) {
		service := args[0]
		user := args[1]

		ok := utils.YesNoPrompt("Are you sure you want to delete the following credential: ["+service+":"+user+"]", true)
		if ok {
			err := keyring.Delete(service, user)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("The following credential has been deleted: [" + service + ":" + user + "]")
		} else {
			fmt.Println("Operation aborted")
		}
	},
}

func init() {
	rootCmd.AddCommand(delCmd)
	getCmd.SetUsageTemplate(fmt.Sprintln(`Usage:
  keyring del <service> <user>`))

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// delCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// delCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
