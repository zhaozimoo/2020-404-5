package main

import (
	"io"
	"net/http"
	"os"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	//POST获取上传的文件并将其保存到磁盘。
	case "POST":
		//解析请求中的multipart form
		err := r.ParseMultipartForm(100000)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		//获取multipart form编号
		m := r.MultipartForm

		//get the *fileheaders
		files := m.File["uploadfile"]
		for i, _ := range files {
			//循环fileheader, 获取实际文件的句柄
			file, err := files[i].Open()
			defer file.Close()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			//create destination file making sure the path is writeable.
			dst, err := os.Create("/data/upload/" + files[i].Filename)
			defer dst.Close()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			//copy the uploaded file to the destination file
			if _, err := io.Copy(dst, file); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

		}

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/upload", uploadHandler)

	//static file handler.
	http.Handle("/staticfile/", http.StripPrefix("/staticfile/", http.FileServer(http.Dir("./staticfile"))))

	//Listen on port 8080
	http.ListenAndServe(":8181", nil)
}