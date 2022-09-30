package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

var (
	serverAddress = flag.String("s", "", "服务端口")
	rs            = flag.String("r", "", "")
)

func main() {
	flag.Parse()
	log.Printf("root==============%s\n", *rs)
	if *serverAddress == "" {
		*serverAddress = os.Getenv("SERVER_ADDRESS")
	}
	roots := strings.Split(*rs, ",")
	for _, root := range roots {
		err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			log.Println(path)
			// 忽略文件夹
			if info.IsDir() {
				return err
			}
			// 忽略html和文本文件
			if strings.HasSuffix(path, ".html") || strings.HasSuffix(path, ".txt") {
				return err
			}
			// 忽略图片
			if strings.HasSuffix(path, ".jpg") || strings.HasSuffix(path, "jpeg") {
				return err
			}
			// 忽略图片
			if strings.HasSuffix(path, ".png") || strings.HasSuffix(path, ".gif") {
				return err
			}
			if err = upload(root, path); err != nil {
				return err
			}
			return os.Remove(path)
		})
		if err != nil {
			log.Printf("filepath walk error:%s", err.Error())
			return
		}
	}
	log.Print("---------------------over--------------------------")
}

// upload 上传文件,适合上传大文件
func upload(root, path string) error {
	r, w := io.Pipe()
	m := multipart.NewWriter(w)
	go func() {
		defer w.Close()
		defer m.Close()
		name, err := filepath.Rel(root, path)
		if err != nil {
			log.Printf("get rel path error:%s", err.Error())
			return
		}
		part, err := m.CreateFormFile("myFile", name)
		if err != nil {
			log.Printf("create form file error:%s", err.Error())
			return
		}
		file, err := os.Open(path)
		if err != nil {
			log.Printf("open file error:%s", err.Error())
			return
		}
		defer file.Close()
		if _, err = io.CopyBuffer(part, file, make([]byte, 1<<20)); err != nil {
			log.Printf("copy file error:%s", err.Error())
			return
		}
	}()
	resp, err := http.Post(fmt.Sprintf(`http://%s/upload`, *serverAddress), m.FormDataContentType(), r)
	if resp != nil && resp.Body != nil {
		defer func() {
			_, _ = io.Copy(ioutil.Discard, resp.Body)
			_ = resp.Body.Close()
		}()
	}
	if err != nil && err != io.EOF {
		log.Println(reflect.TypeOf(err))
		log.Printf("http post error:%s", err.Error())
		return err
	}
	if resp == nil {
		return errors.New("resp is nil")
	}
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("read response error:%s", err.Error())
		return err
	}
	log.Println("response=======", string(bs))
	if resp.StatusCode != 200 {
		return fmt.Errorf("status code :%d", resp.StatusCode)
	}
	return err
}
