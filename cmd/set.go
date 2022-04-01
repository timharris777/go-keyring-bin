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

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Args:  cobra.ExactArgs(2),
	Short: "Create/update a password in the OS credential store",
	Long:  `Create/update a password in the OS credential store`,
	Run: func(cmd *cobra.Command, args []string) {
		service := args[0]
		user := args[1]

		// Check if password exists
		pwd, err := keyring.Get(service, user)
		utils.UNUSED(pwd)
		if err == nil {
			ok := utils.YesNoPrompt("A credential already exists for ["+service+":"+user+"]. Do you want to overwrite it?", true)
			if !ok {
				log.Fatal("Operation aborted")
			}
			password := utils.PasswordPrompt("New password:")

			// create/update credential
			err = keyring.Set(service, user, password)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Password updated succesfully for [" + service + ":" + user + "]")
		} else {
			password := utils.PasswordPrompt("Password:")

			// create/update credential
			err = keyring.Set(service, user, password)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Password set succesfully for [" + service + ":" + user + "]")
		}
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
	getCmd.SetUsageTemplate(fmt.Sprintln(`Usage:
  keyring set <service> <user>`))

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
