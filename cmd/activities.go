/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/giuszeppe/github-activity-go-cli/service"
	"github.com/spf13/cobra"
)

// activitiesCmd represents the activities command
var activitiesCmd = &cobra.Command{
	Use:   "activities",
	Short: "Fetches activities of <username>",
	Long:  `Usage: github-activity fetch activities <username>`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("Argument <username> is required")
		}
		out, err := service.GetActivityForUsername(args[0])

		if err != nil {
			log.Fatal("Error retrieving data")
		}

		if out == "" {
			fmt.Println("No activities for user or user not found.")
			return
		}

		fmt.Println(out)
	},
}

func init() {
	fetchCmd.AddCommand(activitiesCmd)
}
