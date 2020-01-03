package main
import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"fmt"
	"math/rand"
	"time"
	"strings"
)
func CreateRand() string {
	return fmt.Sprintf("%08v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(100000000))
}
var Filename string
// UploadHandler 上传接口
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// 返回上传html页面
		data, err := ioutil.ReadFile("./static/view/index.html")
		if err != nil {
			log.Printf("UploadHandler -> {%s}", err)
			_, _ = io.WriteString(w, "服务器错误")
			return
		}
		_, _ = io.WriteString(w, string(data))
	} else if r.Method == "POST" {
		// 接收文件流
		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			log.Printf("UploadHandler: 上传文件出错 -> {%s}", err)
			_, _ = io.WriteString(w, "上传文件出错！")
			return
		}
		defer file.Close()

		headerByte, _ := json.Marshal(fileHeader.Header)
		log.Printf("当前文件：Filename - >{%s}, Size -> {%v}, FileHeader -> {%s}", fileHeader.Filename, fileHeader.Size, string(headerByte))

		newFile, err := os.Create("/data/resourcecenter/" +"rec"+CreateRand() + fileHeader.Filename)
		if err != nil {
			log.Printf("UploadHandler: 创建文件失败！")
			_, _ = io.WriteString(w, "服务器错误！")
			return
		}
		defer newFile.Close()

		// 复制文件到目标目录
		_, errCopy := io.Copy(newFile, file)
		if errCopy != nil {
			log.Printf("UploadHandler: 文件复制失败！ -> {%s}", err)
			_, _ = io.WriteString(w, "服务器错误！")
			return
		}
		Filename=newFile.Name()
		// 成功响应

		// 重定向到这个请求
		http.Redirect(w, r, "/file/upload/suc", http.StatusFound)

	}
}

// UploadSuccessHandler 响应
func UploadSuccessHandler(w http.ResponseWriter, r *http.Request) {
	comma := strings.Index(Filename, "/")
	pos := strings.Index(Filename[comma:], "rec")
	_, _ = io.WriteString(w, "http://www.icity24.xyz/rc/resourcecenter/"+Filename[comma+pos:])
}


func main() {

	//  http.HandleFunc 指定路由规则

	// 上传文件
	http.HandleFunc("/file/upload", UploadHandler)
	// 上传成功响应，/file/upload 请求处理成功后，会重定向到这个该请求
	http.HandleFunc("/file/upload/suc", UploadSuccessHandler)


	err := http.ListenAndServe(":8980", nil)
	if err != nil {
		fmt.Printf("启动服务器失败: %s", err.Error())
	}
}