package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	flagPath string
)

func init() {
	flag.StringVar(&flagPath, "path", "./", "path, eg: -path ./")
	// flag.
}

func main() {
	flag.Parse()
	walkDir(flagPath, []string{"IMG", "CSS", "JS", "Top.png"}, func(s string) string {
		return strings.ToLower(s)
	})
}

func walkDir(dir string, search []string, transform func(string) string) {
	dirNum, successNum := make([]string, 0), make([]string, 0)
	log.Printf("搜索目录 %s", dir)
	target := make(map[string]bool)
	for _, v := range search {
		target[v] = true
	}
	// 遍历目录
	filepath.Walk(dir, func(filename string, fi os.FileInfo, err error) error {
		log.Printf("文件遍历 %s", fi.Name())

		if err != nil {
			log.Printf("目录遍历失败 %v", err)
			return nil
		}
		fi.IsDir()
		// 忽略目录
		if fi.IsDir() {
			dirNum = append(dirNum, fi.Name())
			if _, ok := target[fi.Name()]; ok {
				t := transform(fi.Name())
				new := strings.Replace(filename, fi.Name(), t, 1)
				if err := os.Rename(filename, new); err != nil {
					log.Printf("目录 %s 修改 %s 失败", filename, new)
				}
				log.Printf("成功将 %s 替换成 %s \n", filename, new)
				successNum = append(successNum, filename)
			}
			return nil
		}
		return nil
	})
	log.Printf("查询到目录：%d 个，修改成功 %d 个", len(dirNum), len(successNum))
}

func isDir(path string) (bool, error) {
	_, _err := os.Stat(path)
	if _err == nil {
		return true, nil
	}
	if os.IsNotExist(_err) {
		return false, nil
	}
	return false, _err
}

// 创建文件夹
func createDir(path string) {
	_exist, _err := isDir(path)
	if _err != nil {
		fmt.Printf("获取文件夹异常 -> %v\n", _err)
		return
	}
	if _exist {
		fmt.Println("文件夹已存在！")
	} else {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			fmt.Printf("创建目录异常 -> %v\n", err)
		} else {
			fmt.Println("创建成功!")
		}
	}
}

// 删除文件
func removeFile(path string) error {
	_err := os.Remove(path)
	return _err
}

// 删除文件夹
func removeDir(path string) error {
	_err := os.RemoveAll(path)
	return _err
}
