package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func header(w http.ResponseWriter, r *http.Request) {
	for key, val := range r.Header {
		_, _ = fmt.Fprintf(w, "%s: %s\n", key, val[0])
	}

	//fmt.Println(r.Header.Get("connection"))
	//fmt.Fprint(w, "------------")
	//fmt.Fprint(w, r.ContentLength)
}

func body(w http.ResponseWriter, r *http.Request) {
	length := r.ContentLength
	body := make([]byte, length)
	_, _ = r.Body.Read(body)
	_, _ = fmt.Fprintf(w, "%s\n", string(body))
}

func index(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprint(w, `<html>
	<head>
	<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
	<title>Go Web Programming</title>
	</head>
	<body>
	<form action=http://127.0.0.1:9000/process?name=xiaofang&boyfriend=longshuai
	method="post" enctype="multipart/form-data">
	<input type="text" name="name" value="longshuai"/>
	<input type="text" name="age" value="23"/>
	<input type="file" name="file_to_upload">
	<input type="submit"/>
	</form>
	</body>
	</html>`)
}

func form(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024)
	//fmt.Fprint(w, r.Form)
	//fmt.Fprint(w, r.PostForm)
	//fmt.Fprint(w, r.MultipartForm)
	fileHeader := r.MultipartForm.File["file_to_upload"][0]
	file, err := fileHeader.Open()
	if err == nil {
		dataFormFile, err := ioutil.ReadAll(file)
		if err == nil{
			fmt.Fprintf(w, "%s\n", dataFormFile)
		}
	}
	fmt.Fprintf(w, "%s\n",r.MultipartForm)
}

func main() {
	server := http.Server{
		Addr: ":9000",
	}

	http.HandleFunc("/headers", header)
	http.HandleFunc("/body", body)
	http.HandleFunc("/process", form)
	http.HandleFunc("/", index)
	_ = server.ListenAndServe()
}
