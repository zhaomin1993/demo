package unzip

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"

	"github.com/axgle/mahonia"
)

// Unzip 解压zip压缩包到指定文件夹下
func Unzip(zipFile string, destDir string) error {
	zipReader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer zipReader.Close()
	var decodeName string
	for _, f := range zipReader.File {
		if err = func(f *zip.File) error {
			if f.NonUTF8 {
				//如果标致位是0  则是默认的本地编码   默认为gbk
				dec := mahonia.NewDecoder("gbk")
				decodeName = dec.ConvertString(f.Name)
			} else {
				//如果标志为是 1 << 11也就是 2048  则是utf-8编码
				decodeName = f.Name
			}
			fpath := filepath.Join(destDir, decodeName)
			if f.FileInfo().IsDir() {
				if err = os.MkdirAll(fpath, os.ModePerm); err != nil {
					return err
				}
			} else {
				if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
					return err
				}

				inFile, err := f.Open()
				if err != nil {
					return err
				}
				defer inFile.Close()

				outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
				if err != nil {
					return err
				}
				defer outFile.Close()

				_, err = io.Copy(outFile, inFile)
				if err != nil {
					return err
				}
			}
			return nil
		}(f); err != nil {
			return err
		}
	}
	return nil
}
