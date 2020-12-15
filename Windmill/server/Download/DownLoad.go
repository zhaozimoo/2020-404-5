package Download
import (
	"log"
	"net/http"
	"os"
	"io"
	"strconv"
)


func HanderGetFile(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r.Body!=nil{
			r.Body.Close()
		}
	}()
	r.ParseForm()  //解析参数，默认是不会解析的
	log.Println("Recv:", r.RemoteAddr)
	pwd,_ := os.Getwd()
	des:= pwd + string(os.PathSeparator) + r.URL.Path[1:len(r.URL.Path)]
	//fmt.Print(r.URL.Path)
	desStat,err := os.Stat(des)
	if err != nil{
		log.Println("File Not Exit", des)
		http.NotFoundHandler().ServeHTTP(w, r)
	}else if(desStat.IsDir()){
		log.Println("File Is Dir", des)
		http.NotFoundHandler().ServeHTTP(w, r)
	}else{
		fileData, err := os.Open(des)
		defer fileData.Close()
		//fileData, err := ioutil.ReadFile(des)
		//fileData,err:=io.Reader(des)
		if err != nil{
			log.Println("Read File Err:", err.Error())
		}else{
			log.Println("Send File:", des)
			//f, err := os.Create(des[25:len(des)])
			////f, err := os.Create("test.png")
			//if err != nil {
			//	panic(err)
			//}
			size:=desStat.Size()
			sizes := strconv.FormatInt(size,10)
			w.Header().Add("Content-Type","application/octet-stream")
			w.Header().Add("Content-Length",sizes)
			io.Copy(w, fileData)
		}
	}
}

