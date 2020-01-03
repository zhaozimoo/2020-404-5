package main
import (
	"fmt"
	"bytes"
	"mime/multipart"
	"os"
	"path/filepath"
	"io"
	"net/http"
	"io/ioutil"
)

func main() {

	url := "http://zwfw.nmg.gov.cn/app/resourcecenterup/file/upload"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, errFile1 := os.Open("/data/test1111.jpg")
	defer file.Close()
	part1,
	errFile1 := writer.CreateFormFile("file",filepath.Base("/data/test1111.jpg"))
	_, errFile1 = io.Copy(part1, file)
	if errFile1 !=nil {

		fmt.Println(errFile1)
	}
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
	}


	client := &http.Client {
	}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "multipart/form-data****")

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}