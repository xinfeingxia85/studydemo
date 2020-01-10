package main

import (
	"github.com/spf13/cobra"
	"io/ioutil"
	"os/exec"
)

func init() {
	RootCmd.AddCommand(stopCmd)
}

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop Gonone",
	Run: func(cmd *cobra.Command, args []string) {
		strb, _ := ioutil.ReadFile("gonone.lock")
		command := exec.Command("kill", string(strb))
		command.Start()
		println("gonone stop")
	},
}
