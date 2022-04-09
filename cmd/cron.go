/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"ayush-deliveroo/internal/service"

	"github.com/spf13/cobra"
)

var cronCmd = &cobra.Command{
	Use: "cron",

	Run: func(cmd *cobra.Command, args []string) {
		service.Cron(args)
	},
}

func init() {
	rootCmd.AddCommand(cronCmd)
}
