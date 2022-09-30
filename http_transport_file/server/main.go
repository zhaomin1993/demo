package main

import (
	"bufio"
	"crypto/tls"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/acme/autocert"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	//r.MaxMultipartMemory = 6 << 30
	r.Use(gin.Recovery(), gin.Logger())
	r.POST("/upload", func(context *gin.Context) {
		fh, err := context.FormFile("myFile")
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"code:": 400, "msg": "get file error:" + err.Error()})
			return
		}
		formFile, openErr := fh.Open()
		if openErr != nil {
			context.JSON(http.StatusBadRequest, gin.H{"code:": 400, "msg": "open form file error:" + openErr.Error()})
			return
		}
		defer formFile.Close()
		filePath := filepath.Join(`/go/srv/data`, fh.Filename)
		fileDir := filepath.Dir(filePath)
		if !Exist(fileDir) {
			if err = os.MkdirAll(fileDir, 0664); err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"code:": 400, "msg": "create dir error:" + err.Error()})
				return
			}
		}
		newFile, createErr := os.Create(filePath)
		if createErr != nil {
			context.JSON(http.StatusBadRequest, gin.H{"code:": 500, "msg": "create new file error:" + createErr.Error()})
			return
		}
		defer newFile.Close()
		wt := bufio.NewWriterSize(newFile, 1<<20)
		if _, err = io.CopyBuffer(wt, formFile, make([]byte, 1<<20)); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"code:": 500, "msg": "io copy file error:" + err.Error()})
			return
		}
		if err = wt.Flush(); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"code:": 500, "msg": "flush file error:" + err.Error()})
			return
		}
		context.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success"})
	})
	m := &autocert.Manager{
		Cache:  autocert.DirCache("secret-dir"),
		Prompt: autocert.AcceptTOS,
		// HostPolicy: autocert.HostWhitelist("example.org"),
	}
	server := &http.Server{
		Addr:           "0.0.0.0:9999",
		Handler:        r,
		ReadTimeout:    time.Hour,
		WriteTimeout:   time.Second * 10,
		MaxHeaderBytes: 1 << 19,
		IdleTimeout:    time.Hour,
		TLSConfig:      &tls.Config{GetCertificate: m.GetCertificate},
	}
	serErr := server.ListenAndServe()
	if serErr != nil {
		log.Printf("Listen And Serve Error:%s", serErr.Error())
	}
}

func Exist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
