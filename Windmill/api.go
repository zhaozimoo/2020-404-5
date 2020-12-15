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

	url := "http://www.icity24.xyz/rc/fileserver/file/upload"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, errFile1 := os.Open("/C:/Users/zhaozimolc/Pictures/Camera Roll/登录提示2x.png")
	defer file.Close()
	part1,
	errFile1 := writer.CreateFormFile("file",filepath.Base("/C:/Users/zhaozimolc/Pictures/Camera Roll/登录提示2x.png"))
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