package Upload

import (
"io"
"net/http"
"os"
"log"
"fmt"
"time"
"strconv"
"strings"
"path"
"encoding/json"
"Windmill/server/RandomStrings"
"Windmill/server/config"
)

type wWriteBody struct {
	Code string `json:"code"`
	Msg string `json:"message"`
	Data interface{} `json:"data"`
}

func returnRes(data interface{},w http.ResponseWriter){
	var writeBody wWriteBody
	writeBody.Code="0000"
	writeBody.Msg="业务正常处理完毕"
	writeBody.Data=data
	writeDate,_:=json.Marshal(writeBody)
	w.Write(writeDate)
}

func returnErr(w http.ResponseWriter){
	var writeBody wWriteBody
	writeBody.Code="0001"
	writeBody.Msg="业务处理异常"
	writeDate,_:=json.Marshal(writeBody)
	w.Write(writeDate)
}

var filename string
var r1 string
var r2 string
var r3 string

func JuploadHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	//POST获取上传的文件并将其保存到磁盘。
	case "POST":
		//解析请求中的multipart form
		err := r.ParseMultipartForm(20000000)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		//获取multipart form编号
		m := r.MultipartForm
		code:=r.FormValue("code")
		fmt.Print(code)
		files := m.File["file"]
		if err != nil {
			log.Printf("UploadHandler: 上传文件出错 -> {%s}", err)
			_, _ = io.WriteString(w, "上传文件出错！")
			returnErr(w)
			return
		}

		for i, _ := range files {
			//循环fileheader, 获取实际文件的句柄
			file, err := files[i].Open()
			defer file.Close()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			fullFilename := files[i].Filename
			fmt.Println("fullFilename =", fullFilename)
			var filenameWithSuffix string
			filenameWithSuffix = path.Base(fullFilename) //获取文件名带后缀
			fmt.Println("filenameWithSuffix =", filenameWithSuffix)
			var fileSuffix string
			fileSuffix = path.Ext(filenameWithSuffix) //获取文件后缀
			fmt.Println("fileSuffix =", fileSuffix)
			var filenameOnly string
			filenameOnly = strings.TrimSuffix(filenameWithSuffix, fileSuffix)//获取文件名
			fmt.Println("filenameOnly =", filenameOnly)
			r1 := strconv.Itoa(time.Now().Year())
			r2=time.Now().Month().String()
			r3 := strconv.Itoa(time.Now().Day())
			cgdirExist(r1,r2,r3)
			fmt.Println("111111111111111111")
			dst, err := os.Create("/clouddata/"+r1+"/"+r2+"/" +r3+"/"+"rec"+RandomStrings.GetRandomStringscg(32)+ fileSuffix)
			fmt.Print(dst.Name())

			fmt.Print(files[i].Filename)
			if err != nil {
				log.Printf("UploadHandler: 创建文件失败！")
				_, _ = io.WriteString(w, "服务器错误！")
				return
			}
			defer dst.Close()

			_, errCopy := io.Copy(dst, file)
			if errCopy != nil {
				log.Printf("UploadHandler: 文件复制失败！ -> {%s}", err)
				_, _ = io.WriteString(w, "服务器错误！")
				return
			}
			filename=dst.Name()
			if code=="ccwork" {
				comma := strings.Index(filename, "/")
				pos := strings.Index(filename[comma:], "rec")
				data:=make(map[string]string)
				//data["fileName"]=fmt.Sprintf(filename[comma+pos:])
				data["download_url"]=fmt.Sprintf(config.GetConfig().Return.DownloadUrl+r1+"/"+r2+"/" +r3+"/"+filename[comma+pos:])
				returnRes(data,w)
			}
		}

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}


func SuploadHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	//POST获取上传的文件并将其保存到磁盘。
	case "POST":
		//解析请求中的multipart form
		err := r.ParseMultipartForm(20000000)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		//获取multipart form编号
		m := r.MultipartForm
		files := m.File["file"]
		if err != nil {
			log.Printf("UploadHandler: 上传文件出错 -> {%s}", err)
			_, _ = io.WriteString(w, "上传文件出错！")
			returnErr(w)
			return
		}

		for i, _ := range files {
			//循环fileheader, 获取实际文件的句柄
			file, err := files[i].Open()
			defer file.Close()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			fullFilename := files[i].Filename
			fmt.Println("fullFilename =", fullFilename)
			var filenameWithSuffix string
			filenameWithSuffix = path.Base(fullFilename) //获取文件名带后缀
			fmt.Println("filenameWithSuffix =", filenameWithSuffix)
			var fileSuffix string
			fileSuffix = path.Ext(filenameWithSuffix) //获取文件后缀
			fmt.Println("fileSuffix =", fileSuffix)
			var filenameOnly string
			filenameOnly = strings.TrimSuffix(filenameWithSuffix, fileSuffix)//获取文件名
			fmt.Println("filenameOnly =", filenameOnly)
			r1 := strconv.Itoa(time.Now().Year())
			r2=time.Now().Month().String()
			r3 := strconv.Itoa(time.Now().Day())
			cgdirExist(r1,r2,r3)
			fmt.Println("111111111111111111")
			dst, err := os.Create("/clouddata/"+r1+"/"+r2+"/" +r3+"/"+"rec"+RandomStrings.GetRandomStringscg(32)+ fileSuffix)
			fmt.Print(dst.Name())

			fmt.Print(files[i].Filename)
			if err != nil {
				log.Printf("UploadHandler: 创建文件失败！")
				_, _ = io.WriteString(w, "服务器错误！")
				return
			}
			defer dst.Close()

			_, errCopy := io.Copy(dst, file)
			if errCopy != nil {
				log.Printf("UploadHandler: 文件复制失败！ -> {%s}", err)
				_, _ = io.WriteString(w, "服务器错误！")
				return
			}
			filename=dst.Name()
			comma := strings.Index(filename, "/")
			pos := strings.Index(filename[comma:], "rec")
			w.Write([]byte(config.GetConfig().Return.DownloadUrl+r1+"/"+r2+"/" +r3+"/"+filename[comma+pos:]))
		}

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func cgdirExist (dir1 string,dir2 string,dir3 string ) {
	sta,err := os.Stat("/clouddata/"+dir1+"/"+dir2+"/" +dir3)
	if err != nil {
		fmt.Println("stat temp dir error,maybe is not exist, maybe not")
		if os.IsNotExist(err) {
			fmt.Println("temp dir is not exist")
			err := os.MkdirAll("/clouddata/"+dir1+"/"+dir2+"/" +dir3, os.ModePerm)
			if err != nil {
				fmt.Printf("mkdir failed![%v]\n", err)
			}
			os.Chmod("/clouddata/"+dir1+"/"+dir2+"/" +dir3, 0777)
			return
		}

		fmt.Println("stat file error")
		fmt.Println(sta)
		return
	}
	fmt.Println(sta)
	fmt.Println("temp_dir is exist")
}


