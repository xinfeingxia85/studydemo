package main

import (
	"fangkm.demo/cmd/api"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
)

var excuteCount = 0

const LOCK_FILE_NAME = "gonone.lock"

func init() {
	var daemon bool

	startCmd := &cobra.Command{
		Use:   "start",
		Short: "start gonone",
		Run: func(cmd *cobra.Command, args []string) {
			//check if the
			if daemon {
				startStatus := checkServiceIsExist(LOCK_FILE_NAME)
				if startStatus {
					fmt.Print("[Error] gonone start is already running...")
					os.Exit(1)
				}
				command := exec.Command("./gonone", "start")
				_ = command.Start()
				fmt.Printf("gonone starts, [PID] %d running...\n", command.Process.Pid)
				_ = ioutil.WriteFile(LOCK_FILE_NAME, []byte(fmt.Sprintf("%d", command.Process.Pid)), 0666)
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
		Addr:    ":9004",
		Handler: api.NewApiMux(),
	}

	//http.HandlerFunc()
	//http.HandleFunc("/", Index)
	//http.Ha
	_ = server.ListenAndServe()
}

//check whether the port is already existing
func checkServiceIsExist(path string) bool {
	isExist := FileExist(path)
	if !isExist {
		return false
	}

	//check whether ther content of one file  is already existing
	data, _ := ioutil.ReadFile(LOCK_FILE_NAME)
	if len(data) > 0 {
		return true
	}

	return false
}

func FileExist(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}

//func Index(w http.ResponseWriter, r *http.Request) {
//	var rStr = "http sucess" + strconv.Itoa(excuteCount)
//	w.Write([]byte(rStr))
//}
