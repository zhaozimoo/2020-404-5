package main

import (
	"net/http"
    "Windmill/server/Upload"
	"Windmill/server/config"
	"Windmill/server/Download"
)
func main() {
	//上传入口
	http.HandleFunc("/file/upload", Upload.JuploadHandler)
	http.HandleFunc("/upload", Upload.SuploadHandler)
	http.HandleFunc("/", Download.HanderGetFile)
	http.ListenAndServe(config.GetConfig().Return.UploadServerPort, nil)

   //下载入口
	//port := flag.Int("p", config.GetConfig().Return.DownServerPort, "Set The Http Port")
	//flag.Parse()
	//pwd,_ := os.Getwd()
	//log.Printf("Listen On Port:%d pwd:%s\n", *port, pwd)
   //
	//http.HandleFunc("/", Download.HanderGetFile)
	//err := http.ListenAndServe(":" + strconv.Itoa(*port), nil)
	//if nil != err{
	//	log.Fatalln("Get Dir Err", err.Error())
	//}
}