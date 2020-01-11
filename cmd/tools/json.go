package tools

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

func MarshalJson(w http.ResponseWriter, playload interface{}) {
	data, err := json.Marshal(playload)
	ThrowError(err)
	w.Header().Set("Content-type", "application/json")
	_, _ = w.Write(data)
}

func UnMarshalJson(req *http.Request, playload interface{}) {
	data, err := ioutil.ReadAll(req.Body)
	err = errors.New("test")
	ThrowError(err)
	_ = json.Unmarshal(data, playload)
}
