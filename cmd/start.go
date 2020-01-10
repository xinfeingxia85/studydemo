package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strconv"
)

var excuteCount = 0

func init() {
	var daemon bool

	startCmd := &cobra.Command{
		Use:   "start",
		Short: "start gonone",
		Run: func(cmd *cobra.Command, args []string) {
			if daemon {
				command := exec.Command("./gonone", "start")
				command.Start()
				fmt.Printf("gonone starts, [PID] %d running...\n", command.Process.Pid)
				ioutil.WriteFile("gonone.lock", []byte(fmt.Sprintf("%d", command.Process.Pid)), 0666)
				daemon = false
				os.Exit(0)
			} else {
				excuteCount = excuteCount + 1
				fmt.Println("no daemon!!")
			}
			fmt.Print("test...")
			startHttp()
		},
	}

	startCmd.Flags().BoolVarP(&daemon, "daemon", "d", false, "is daemon?")
	RootCmd.AddCommand(startCmd)
}

func startHttp() {
	server := http.Server{
		Addr: ":9004",
	}

	//http.HandlerFunc()
	http.HandleFunc("/", Index)
	server.ListenAndServe()
}

func Index(w http.ResponseWriter, r *http.Request) {
	var rStr = "http sucess" + strconv.Itoa(excuteCount)
	w.Write([]byte(rStr))
}
