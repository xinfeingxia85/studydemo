package main

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "gonone",
	Run: func(cmd *cobra.Command, args []string) {
		print("gonne is my ai friend")
	},
}
