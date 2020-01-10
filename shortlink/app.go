package shortlink

import (
	"encoding/json"
	"log"
	"net/http"
)

type App struct {
	Server *http.Server
}

type Request struct {
	URL             string `json:"url"`
	ExpireInMinuter int32  `json:"expire_in_minuter"`
}

type Response struct {
	ShortLink string `json:"short_link"`
}

//create app
func NewApp(addr string, options ...func(server *http.Server)) *App {
	srv := &http.Server{
		Addr: addr,
	}

	for _, option := range options {
		option(srv)
	}

	return &App{srv}
}

//init router
func (app *App) Router() {
	//init mux
	mux := http.NewServeMux()
	mux.HandleFunc("/createlink/", CreateLink)
	app.Server.Handler = mux
}

//start http server
func (app *App) Run() {
	app.Router()
	log.Fatal(app.Server.ListenAndServe())
}

//handler method
func CreateLink(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	log.Println(r.FormValue("url"))
	bytes, _ := json.Marshal(Response{
		ShortLink: "http://fangk.com/djkfk/",
	})
	_, _ = w.Write(bytes)
}