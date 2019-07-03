package main

import (
	"archive/tar"
	"time"
	"io/ioutil"
	"log"
	"bytes"
	"fmt"
	"os"
	"io"
)
// 文件类型
const (
	TypeReg           = tar.TypeReg    		    // 普通文件
	//TypeRegA          = tar.TypeRegA 		    // 普通文件 Deprecated: Use TypeReg instead.
	TypeLink          = tar.TypeLink    		// 硬链接(hard link)
	TypeSymlink       = tar.TypeSymlink   		// 符号链接
	TypeChar          = tar.TypeChar    		// 字符设备节点
	TypeBlock         = tar.TypeBlock    		// 块设备节点
	TypeDir           = tar.TypeDir    			// 目录
	TypeFifo          = tar.TypeFifo    	    // fifo节点
	TypeCont          = tar.TypeCont    		// 保留
	TypeXHeader       = tar.TypeXHeader   		// 扩展标题
	TypeXGlobalHeader = tar.TypeXGlobalHeader   // 全局扩展标题
	TypeGNULongName   = tar.TypeGNULongName     // 下一个文件名称很长
	TypeGNULongLink   = tar.TypeGNULongLink     // 接下来将文件符号链接到一个带有长名字的文件
	TypeGNUSparse     = tar.TypeGNUSparse    	// 稀疏文件(sparse file)
 )
const FilePath  = "/Users/artist/go/src/gostudy/standard_library/archive/zip/tartest/tarexample.zip"

func main() {
	w := write()
	if err := ioutil.WriteFile(FilePath, w.Bytes(), os.ModePerm); err != nil{
		log.Fatal(err)
	}
	read(FilePath)
}
func write()bytes.Buffer  {
	var buf bytes.Buffer
	var tw = tar.NewWriter(&buf)
	var files = []struct{
		Name, Body string
	}{
		{"readme.txt", "This archive contains some text files."},
		{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		{"todo.txt", "Get animal handling license."},
	}
	for _, file := range files{
		hdr := &tar.Header{
			Name:       file.Name,             // 记录头域的文件名
			Mode:       0600,                  // 权限和模式位
			Uid:        0,                     // 所有者的用户ID
			Gid:        0,                     // 所有者的组ID
			Size:       int64(len(file.Body)), // 字节数（长度）
			ModTime:    time.Now(),            // 修改时间
			Typeflag:   TypeReg,               // 文件类型
			Linkname:   "",                    // 链接的目标名
			Uname:      "",                    // 所有者的用户名
			Gname:      "",                    // 所有者的组名
			Devmajor:   0,                     // 字符设备或块设备的major number
			Devminor:   0,                     // 字符设备或块设备的minor number
			AccessTime: time.Now(),            // 访问时间
			ChangeTime: time.Now(),            // 状态改变时间
		}
		if err := tw.WriteHeader(hdr); err != nil {
			log.Fatal(err)
		}
		if _, err := tw.Write([]byte(file.Body)); err != nil {
			log.Fatal(err)
		}
	}
	if err := tw.Close(); err != nil {
		log.Fatal(err)
	}
	return buf
}
func read(path string)  {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	// 初始化一个reader去读取tar内容
	r := tar.NewReader(f)
	// 循环读取多个文件内容
	for{
		hdr, err := r.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Contents of %s:\n", hdr.Name)
		content, err := ioutil.ReadAll(r)
		fmt.Printf("%s\n", content)
	}
}