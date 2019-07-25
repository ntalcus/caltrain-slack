package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "ctc",
	Short: "caltrain-cli is a CLI to access transit info",
	Long: `A tool to access information provided by 511.org
				  for all your scheduling needs.
				  Created by Noah Alcus.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
