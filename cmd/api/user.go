package api

import (
	"fangkm.demo/cmd/model"
	"fangkm.demo/cmd/tools"
	"github.com/gorilla/mux"
	"net/http"
)

func initUserApi(s *mux.Router) {
	r := s.PathPrefix("/user").Subrouter()
	r.HandleFunc("/list", UserListHandle)
	r.HandleFunc("/add", UserAddHandle)
	r.HandleFunc("/login", UserLoginHandle)
}

func UserListHandle(w http.ResponseWriter, r *http.Request) {
	resp := model.Resp{Code: "0", Message: "Sucess"}
	resp.Data = []model.User{
		{Name: "fangkaiming"},
		{Name: "fangzhenyu"},
	}

	//data, err := json.Marshal(&resp)
	//if err != nil {
	//	panic(err)
	//}
	//w.Header().Set("Content-type", "application/json")
	//_, _ = w.Write(data)
	tools.MarshalJson(w, &resp)
}

func UserAddHandle(w http.ResponseWriter, r *http.Request) {
	resp := model.Resp{
		Code:    "0",
		Message: "sucess",
	}

	tools.MarshalJson(w, &resp)

	//w.Header().Set("Content-type", "application/json")
	//data, err := json.Marshal(&resp)
	//if err != nil {
	//	panic(err)
	//}
	//
	//_, _ = w.Write(data)
}

func UserLoginHandle(w http.ResponseWriter, r *http.Request) {
	user := &model.User{}
	//err := json.NewDecoder(r.Body).Decode(&user)
	//if err != nil {
	//	//panic(err)
	//	http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	//	return
	//}
	defer tools.CatchError(w, r)
	tools.UnMarshalJson(r, &user)

	resp := model.Resp{
		Code:    "1001",
		Message: "账号或者密码错误!",
	}

	if user.Name == "fangkm" && user.Password == "123" {
		resp = model.Resp{
			Code:    "0",
			Message: "sucess",
		}
	}

	tools.MarshalJson(w, &resp)

	//data, err := json.Marshal(&resp)
	//if err != nil {
	//	panic(err)
	//}
	//
	//_, _ = w.Write(data)
}
