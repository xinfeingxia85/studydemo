package main

import (
	"github.com/billcobbler/casbin-redis-watcher/v2"
	"github.com/casbin/casbin/v2"
	"log"
)

func main() {
	w, _ := rediswatcher.NewWatcher("127.0.0.1:6379")


	e, err := casbin.NewEnforcer("config/model.conf", "config/policy.csv")
	if err != nil {
		log.Println(err)
	}

	_ = e.SetWatcher(w)
	_ = w.SetUpdateCallback(updateMsg)
	_ = e.SavePolicy()
}

func updateMsg(msg string) {
	log.Println(msg)
}
