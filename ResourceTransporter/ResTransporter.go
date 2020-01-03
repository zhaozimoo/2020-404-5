package main
import (
	"Logger/logger"
	"net/http"
	"ResourceTransporter/Transporter"
)

//func init(){
//     parseJob.C=parseJob.GetCron()
//	parseJob.C.Start()
//}

//func Startjob(){
//	parseJob.DownloadJosn()
//}

func main() {

	logger.InitLogConfig(logger.DEBUG, true)
	http.HandleFunc("/start",func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		env:=r.PostFormValue("env")
		projectname:=r.PostFormValue("project_name")
		res,err:=Transporter.RunShell(env,projectname)
		if err!=nil{
			logger.Error(err)
			return
		}
		w.Write([]byte(res))
		//data:=make(map[string]string)
		//data["result"]=fmt.Sprintf("添加数据%d条",res)
		//ReturnRes(data,w)

	})
	if err:=http.ListenAndServe(":11009", nil);err!=nil{
		logger.Error("ListenAndServe err",err)     }
}